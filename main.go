package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Meetup API Key
const apiKey = "77785333e3c2f604b62487138582a69"

// Topics to pull Events from
var topics = []string{
	"sysadmin",
	"devops",
	"infrastructure-as-code",
	"docker",
	"opensource",
	"cloud-computing",
	"automation",
	"kubernetes",
	"containers",
	"iot-hacking",
	"internet-of-things",
	"sensors",
	"amazon-web-services",
	"high-scalability-computing",
	"saas-software-as-a-service",
}

// Areas to pull meetups from
var calls = []*call{
	NewCall("SF", callUrl("37.7749", "-122.4194")),
	NewCall("UK", callUrl("51.5074", "-0.1278")),
	NewCall("NYC", callUrl("40.7128", "-74.0059")),
	NewCall("DEN", callUrl("39.7392", "-104.9903")),
	NewCall("BOS", callUrl("42.3601", "-71.0589")),
}

// Parses the API Response
type OpenEvents struct {
	Results []struct {
		UtcOffset      int     `json:"utc_offset"`
		RsvpLimit      int     `json:"rsvp_limit,omitempty"`
		Headcount      int     `json:"headcount"`
		Distance       float64 `json:"distance"`
		Visibility     string  `json:"visibility"`
		WaitlistCount  int     `json:"waitlist_count"`
		Created        int64   `json:"created"`
		MaybeRsvpCount int     `json:"maybe_rsvp_count"`
		Description    string  `json:"description"`
		EventURL       string  `json:"event_url"`
		YesRsvpCount   int     `json:"yes_rsvp_count"`
		Duration       int     `json:"duration,omitempty"`
		Name           string  `json:"name"`
		ID             string  `json:"id"`
		Time           int64   `json:"time"`
		Updated        int64   `json:"updated"`
		Group          struct {
			JoinMode string  `json:"join_mode"`
			Created  int64   `json:"created"`
			Name     string  `json:"name"`
			GroupLon float64 `json:"group_lon"`
			ID       int     `json:"id"`
			Urlname  string  `json:"urlname"`
			GroupLat float64 `json:"group_lat"`
			Who      string  `json:"who"`
		} `json:"group"`
		Status string `json:"status"`
		Venue  struct {
			Country              string  `json:"country"`
			LocalizedCountryName string  `json:"localized_country_name"`
			City                 string  `json:"city"`
			Address1             string  `json:"address_1"`
			Name                 string  `json:"name"`
			Lon                  float64 `json:"lon"`
			ID                   int     `json:"id"`
			State                string  `json:"state"`
			Lat                  float64 `json:"lat"`
			Repinned             bool    `json:"repinned"`
		} `json:"venue,omitempty"`
		HowToFindUs string `json:"how_to_find_us,omitempty"`
		Fee         struct {
			Amount      float64 `json:"amount"`
			Accepts     string  `json:"accepts"`
			Description string  `json:"description"`
			Currency    string  `json:"currency"`
			Label       string  `json:"label"`
			Required    string  `json:"required"`
		} `json:"fee,omitempty"`
	} `json:"results"`
	Meta struct {
		Next        string `json:"next"`
		Method      string `json:"method"`
		TotalCount  int    `json:"total_count"`
		Link        string `json:"link"`
		Count       int    `json:"count"`
		City        string
		Description string  `json:"description"`
		Lon         float64 `json:"lon"`
		Title       string  `json:"title"`
		URL         string  `json:"url"`
		ID          string  `json:"id"`
		Updated     int64   `json:"updated"`
		Lat         float64 `json:"lat"`
	} `json:"meta"`
}

// NewCall creates a new call
func NewCall(name, url string) *call {
	return &call{
		name: name,
		url:  url,
	}
}

// callUrl generates the URL for the API Calls
func callUrl(lat, lon string) string {
	template := "https://api.meetup.com/2/open_events?key=%s&topic=%s&status=upcoming&text_format=plain&sign=true&lat=%s&lon=%s"
	return fmt.Sprintf(template, apiKey, strings.Join(topics, ","), lat, lon)
}

// Makes the API Call and returns the results and the next call if there is one
func (c *call) Execute() (OpenEvents, *call) {
	// Tnstantiate return object
	var oe OpenEvents
	var next *call

	fmt.Println(c.url)
	// Prepare request
	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
		return oe, next
	}

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error making request: ", err)
		return oe, next
	}
	fmt.Println("X-RateLimit-Remaining: ", resp.Header.Get("X-RateLimit-Remaining"))
	defer resp.Body.Close()

	// Parse request into
	if err := json.NewDecoder(resp.Body).Decode(&oe); err != nil {
		log.Fatal("Error parsing json: ", err)
	}

	fmt.Println("Records returned: ", oe.Meta.Count)

	// Return parsed request and the next call
	return oe, NewCall(c.name, oe.Meta.Next)
}

// Headers for the CSV file. Change those here!
func headers() []string {
	return []string{"time", "groupName", "eventName", "duration", "eventURL", "rsvpLimit", "yesRSVPCount", "city"}
}

// EventRow Struct
type eventRow struct {
	id             string
	groupName      string
	eventName      string
	duration       time.Duration
	eventURL       string
	maybeRSVPCount int
	rsvpLimit      int
	city           string
	time           time.Time
	yesRSVPCount   int
	description    string
}

// Row formatting for CSV file
func (e eventRow) Strings() []string {
	return []string{e.time.Format(time.RFC3339), e.groupName, e.eventName, e.duration.String(), e.eventURL, fmt.Sprint(e.rsvpLimit), fmt.Sprint(e.yesRSVPCount), e.city}
}

// API Call struct
type call struct {
	name string
	url  string
}

func main() {

	// Initialize variables
	var oe []OpenEvents
	var rows []eventRow

	// Make calls, store unmarshalled structs for processing, follow next links
	for _, c := range calls {
		for {
			if c.url == "" {
				break
			}
			response, next := c.Execute()
			response.Meta.City = next.name
			oe = append(oe, response)
			c = next
		}
		fmt.Println("Responses: ", len(oe))
	}

	// Format things properly
	for _, call := range oe {
		for _, event := range call.Results {
			fee := event.Fee.Required
			switch event {
			case fee == "1":
				continue
			default:

			}
			var eRow = eventRow{
				id:           event.ID,
				groupName:    event.Group.Name,
				city:         call.Meta.City,
				eventName:    event.Name,
				duration:     time.Duration(event.Duration/1000) * time.Second,
				eventURL:     event.EventURL,
				rsvpLimit:    event.RsvpLimit,
				time:         time.Unix(0, event.Time*1000000),
				yesRSVPCount: event.YesRsvpCount,
				description:  event.Description,
			}
			rows = append(rows, eRow)
		}
	}

	// Create CSV file to store results
	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create CSV Writer
	writer := csv.NewWriter(file)

	// Write headers to file
	err = writer.Write(headers())
	if err != nil {
		log.Fatal(err)
		return
	}

	// Write CSV data to file
	for _, row := range rows {
		err = writer.Write(row.Strings())
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	// Clean up!
	writer.Flush()
	file.Close()
}
