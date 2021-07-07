package httptest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/cs-ut-ee/hw3-group-3/pkg/domain"
)

func ApiUrl() string {
	url, success := os.LookupEnv("httpUrl")
	if !success {
		panic("Environment variable 'httpUrl' is not defined")
	}

	return url
}

func TestGetAllPlants(t *testing.T) {

	resp, err := http.Get(ApiUrl() + "/plants")
	if err != nil {
		t.Error("TestGetAllPlants: Problem reading plants via REST.")
		return
	}

	plantsJSON, _ := ioutil.ReadAll(resp.Body)
	plants := []*domain.Plant{}
	json.Unmarshal(plantsJSON, &plants)
	if len(plants) < 2 { // since we have 2 records initially in postgress and it's not going to be chnaged
		t.Error("TestGetAllPlants: there should 2 plants available!")
		return
	}
}

func TestGetPlantPrice(t *testing.T) {

	resp, err := http.Get(ApiUrl() + "/plants/GetPrice/1/2020-03-10/2020-03-12")
	if err != nil {
		t.Error("TestGetPlantPrice: Problem getting plant price via REST.")
		return
	}

	plantPriceJSON, _ := ioutil.ReadAll(resp.Body)
	var plantPrice float64
	json.Unmarshal(plantPriceJSON, &plantPrice)
	if plantPrice != 21.2 { // for 2 days this is the price, this should be in postgres
		t.Error("TestGetPlantPrice: there is something wrong with calculating the price!")
		return
	}
}

func TestIsPlantAvailable(t *testing.T) {

	resp, err := http.Get(ApiUrl() + "/plants/IsPlantAvailable/1/2020-03-11/2020-03-13")
	if err != nil {
		t.Error("TestIsPlantAvailable: Problem testinf if the plant is available via REST.")
		return
	}

	isAvailableJSON, _ := ioutil.ReadAll(resp.Body)
	var isAvailable bool
	json.Unmarshal(isAvailableJSON, &isAvailable)
	if isAvailable { // this item in the postgres has order in this time period (it's created when db is initialized)
		t.Error("TestIsPlantAvailable: the plant should not be available!")
		return
	}
}
