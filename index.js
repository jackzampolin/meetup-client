require('dotenv').config()
const fetch = require('node-fetch')
const {json} = require('micro')
const meetup_api_key = process.env.MEETUP_API_KEY
const _ = require('lodash')

var topics = ["sysadmin", "devops", "infrastructure-as-code", "docker", "opensource", "cloud-computing", "automation", "kubernetes", "containers", "iot-hacking", "internet-of-things", "sensors", "amazon-web-services", "high-scalability-computing", "saas-software-as-a-service"]

var meetupURL = (params) => {
  var enc = Object.keys(params).map(function(key) {
      return key + '=' + params[key];
  }).join('&');
  return `https://api.meetup.com/2/open_events?${enc}`
}

module.exports = async function (req, res) {
  params = {
    "sign": true,
    "key": meetup_api_key,
    "zip": "94105",
    "topics": topics.join(",")
  }
  const response = await fetch(meetupURL(params))
  const resp_json = await response.json()
  console.log(resp_json)
  return "FOO"
}