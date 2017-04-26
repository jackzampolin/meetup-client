# Meetup Client

DOCS: https://www.meetup.com/meetup_api/docs/2/open_events/

This meetup client pulls data about meetups and puts relevant info from each meetup into a csv file. The following topics are desired:

```go
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
```

The following locations are useful to grab meetups from:

```go
var calls = []*call{
	NewCall("SF", callUrl("37.7749", "-122.4194")),
	NewCall("UK", callUrl("51.5074", "-0.1278")),
	NewCall("NYC", callUrl("40.7128", "-74.0059")),
	NewCall("DEN", callUrl("39.7392", "-104.9903")),
	NewCall("BOS", callUrl("42.3601", "-71.0589")),
}
```

To run just `go run main.go`! The file `results.csv` will be created with what you want.