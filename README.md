# getir-go-challange


## Architecture

https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/

---
## Build

`make`

---
## Run Tests

`make test`

---


## API Requests

*TLDR*: `/requests.sh`

#### 1. For the First Endpoint (fetch data from mongodb)
```
curl --request GET \
  --url 'http://localhost:8080/v1/records?startDate=2016-01-26&endDate=2018-02-02&minCount=2700&maxCount=3000' \
  --header 'Content-Type: application/json'
```
---
#### 2. In-memory Endpoints

##### Create configuration
```
curl --request POST \
  --url http://localhost:8080/v1/configurations \
  --header 'Content-Type: application/json' \
  --data '{
	"key": "active-tabs",
	"value": "getir"
}'
```
##### Get configuration

```
curl --request GET \
  --url 'http://localhost:8080/v1/configurations?key=active-tabs' \
  --header 'Content-Type: application/json'
```

