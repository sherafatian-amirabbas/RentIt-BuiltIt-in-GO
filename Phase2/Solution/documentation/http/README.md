# HTTP for BuildIT
Company BuildIT requires HTTP interface for accessing there service. 

# Implementation 
HTTP interface is implemented for accessing three services as follows:
* Get All available Plants
* Get Price of a Plant for a duration
* Check a specific Plant is available or not  

This inteface uses github.com/gorilla/mux for routing, net/http package along with other required package. 

### Get list of all available plants

**Method**   : GET  
**EndPoint** : www.example.com/plants 

**Sample Response:**  
```
[
  {"Id":1,"Name":"eq1","Description":"desc1","PricePerDay":10.5},
  {"Id":2,"Name":"eq2","Description":"desc2","PricePerDay":16.65}
]
```

### Get price of a Plant

**Method**  : GET  
**EndPoint**  : www.example.com/plants/GetPrice/plant_id/start_date/end_date  

**Sample Response:** 
```
{"PlantId":1,"StartDate":"2021-02-02","EndDate":"2021-03-02","PricePerDuration":294}
```

### Check the Plant is available or not

**Method**  : GET  
**EndPoint**  : www.example.com/plants/IsPlantAvailable/plant_id/start_date/end_date  

**Sample Response:** 
```
false
```
