# Meetup Client

Useful links:
- Meetup API Client Docs: https://github.com/jkutianski/meetup-api/wiki
- micro docs: https://github.com/zeit/micro/
- Meetup API Docs: https://www.meetup.com/meetup_api/docs/2/open_events/

This meetup client pulls data about meetups and puts it into a Postgresql database. The following topics are important:

```json
[
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
  "saas-software-as-a-service"
]
```

The following locations are useful to grab meetups from:

```json
[
  {
    "lat": "51.5074",
    "lng": "-0.1278",
    "human": "london"
  },
  {
    "lat": "40.7128",
    "lng": "-74.0059",
    "human": "nyc"
  },
  {
    "lat": "37.7749",
    "lng": "-122.4194",
    "human": "sf"
  },
]
```

Currently this repo can just pull the data it needs and does not do anything with it.