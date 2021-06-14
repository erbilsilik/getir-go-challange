# getir-go-challange

## Architecture

https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/

## Running

`docker-compose up`


## Test

`curl --request GET \
  --url 'http://localhost:8080/v1/records?startDate=2016-01-26&endDate=2018-02-02&minCount=2700&maxCount=3000' \
  --header 'Content-Type: application/json'`


## Further improvements

- Unit/Integration tests
- CI/CD pipeline
