# Websocket
Company ReBuildIT needed websocket interface and because of that there is websocket interface in homework 2

## Implementation 
Rentit application uses github.com/gorilla/websocket module. Because it is better than built in websocket implementation and it also one of the most popular one. Using popular package has some advantages like it has few bugs and security is probably better than some of the other packages. 

Websocket starts with handshake. That during that rentit application upgrades http (**http://**) protocol to websocket protocol (**ws://**).
When it succeeds then you can start sending message to ther server and server can message back. Server can also send messages without client request anything that makes it good for application that needs live data. 

Websocket interface have three main endpoints. 
* Get all plants
* Get plant price
* Is plant available

For that there custom message type is defined. Internally it uses json. \
Message have to have following structure 
```
{
  "resource": "resource-name",
  "params": { "filter-sort-parameter": "value" }
}
```

If you send invalid params then server tries its best to tell you that. 

## Usage
Connect to websocket interface. Url is ws://ip-addr:8080/websocket .
### To get all plants send following message
```
{
  "resource": "plants"
}
```
Sample response 
```
[
  {"Id":1,"Name":"eq1","Description":"desc1","PricePerDay":10.5},
  {"Id":2,"Name":"eq2","Description":"desc2","PricePerDay":16.65}
]
```
### To get plant price send following message
```
{
  "resource": "plant/price",
  "params": { "plantId": "1", "from": "2021-02-02", "to": "2021-03-02"}
}
```
Sample response 
```
{"PlantId":1,"StartDate":"2021-02-02","EndDate":"2021-03-02","PricePerDuration":294}
```

### To check if plant is available on date range send following message
```
{
  "resource": "plant/available",
  "params": { "plantId": "1", "from": "2021-02-02", "to": "2021-03-02"}
}
```
Sample response 
```
true
```

## Testing
Easiest to test websocket interface is to use Firefox **Simple websocket client** plugin \
Our test server url is ws://95.216.188.131:8080/websocket
### Send message 
```
{
  "resource": "plants"
}
```
Sample response 
```
[
  {"Id":1,"Name":"eq1","Description":"desc1","PricePerDay":10.5},
  {"Id":2,"Name":"eq2","Description":"desc2","PricePerDay":16.65}
]
```
