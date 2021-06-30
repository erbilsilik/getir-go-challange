#!/bin/bash

recordsUrl='http://localhost:8080/v1/records?startDate=2016-01-26&endDate=2018-02-02&minCount=2700&maxCount=3000'
PRINTF $"URL: $recordsUrl\n\n"
curl --request GET \
  --url "$recordsUrl" \
  --header 'Content-Type: application/json'
PRINTF $"\n\n"

configurationUrl='http://localhost:8080/v1/configurations'
PRINTF $"URL: $configurationUrl [POST] \n\n"
curl --request POST \
  --url "$configurationUrl" \
  --header 'Content-Type: application/json' \
  --data '{
	"key": "active-tabs",
	"value": "getir"
}'
PRINTF $"\n\n"

PRINTF "URL: $configurationUrl . '?key=active-tabs' [GET]\n\n"

curl --request GET \
  --url "$configurationUrl"'?key=active-tabs' \
  --header 'Content-Type: application/json'
