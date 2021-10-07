# auto-httprequest

auto-httprequest is an util written in go that periodically sends httprequests. 

## Build

The application can be build using `go build`   



## Configuration

The parameters can be set via environment variables

|       VARIABLE|EXAMPLE|Description|
|----------------|-------------------------------|-----------------------------|
|AHTTP_URL|`http://localhost:8080/endpoint`            |Connection URL / Endpoint to be targeted|
|AHTTP_BODY|`{"url":"https://rss.cluster.wenisch.tech/klenkes/events/ical", "filename":"klenkes.ical"}`            |Data that should be attached to the body. For example a JSON            |
|AHTTP_CONTENTTYPE|`application/json`|Content type to be used|
|AHTTP_METHOD|`POST`|Method to be used (GET, PUT, POST, DELETE)|
|AHTTP_WAITTIME|`30s`|Waittime to be used before restarting.|
