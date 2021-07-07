package websockettest

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func ApiUrl() string {
	url, success := os.LookupEnv("websocketUrl")
	if !success {
		panic("Environment variable 'websocketUrl' is not defined")
	}

	return url
}

func TestGetAllPlants(t *testing.T) {
	connection, _, err := websocket.DefaultDialer.Dial(ApiUrl(), nil)
	if err != nil {
		t.Error("Dial:", err)
	}
	defer connection.Close()

	request := "{ \"resource\":\"plants\" }"
	err = connection.WriteMessage(websocket.TextMessage, []byte(request))
	if err != nil {
		t.Error("Send message:", err)
	}

	mt, m, err := connection.ReadMessage()
	if err != nil {
		t.Error("Read message:", err)
	}
	expected := "[{\"Id\":1,\"Name\":\"eq1\",\"Description\":\"desc1\",\"PricePerDay\":10.5},{\"Id\":2,\"Name\":\"eq2\",\"Description\":\"desc2\",\"PricePerDay\":16.65}]"
	if mt != websocket.TextMessage {
		t.Error("Expected message type: " + fmt.Sprint(websocket.TextMessage) + "\nActual: " + fmt.Sprint(mt))
	}

	if string(m) != expected {
		t.Error("Expected: " + string(expected) + "\nActual: " + string(m))
	}

	connection.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(10*time.Second))
}

func TestGetPlantPrice(t *testing.T) {
	connection, _, err := websocket.DefaultDialer.Dial(ApiUrl(), nil)
	if err != nil {
		t.Error("Dial:", err)
	}
	defer connection.Close()

	request := "{ \"resource\":\"plant/available\", \"params\": {\"plantId\":\"1\", \"from\":\"2020-01-01\", \"to\":\"2021-01-01\"} }"
	err = connection.WriteMessage(websocket.TextMessage, []byte(request))
	if err != nil {
		t.Error("Send message:", err)
	}

	mt, m, err := connection.ReadMessage()
	if err != nil {
		t.Error("Read message:", err)
	}
	expected := "true"
	if mt != websocket.TextMessage {
		t.Error("Expected message type: " + fmt.Sprint(websocket.TextMessage) + "\nActual: " + fmt.Sprint(mt))
	}

	if string(m) != expected {
		t.Error("Expected: " + string(expected) + "\nActual: " + string(m))
	}

	connection.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(10*time.Second))
}

func TestGetPlantAvailable(t *testing.T) {
	connection, _, err := websocket.DefaultDialer.Dial(ApiUrl(), nil)
	if err != nil {
		t.Error("Dial:", err)
	}
	defer connection.Close()

	request := "{ \"resource\":\"plant/price\", \"params\": {\"plantId\":\"1\", \"from\":\"2020-01-01\", \"to\":\"2021-01-01\"} }"
	err = connection.WriteMessage(websocket.TextMessage, []byte(request))
	if err != nil {
		t.Error("Send message:", err)
	}

	mt, m, err := connection.ReadMessage()
	if err != nil {
		t.Error("Read message:", err)
	}
	expected := "{\"PlantId\":1,\"StartDate\":\"2020-01-01\",\"EndDate\":\"2021-01-01\",\"PricePerDuration\":3843}"
	if mt != websocket.TextMessage {
		t.Error("Expected message type: " + fmt.Sprint(websocket.TextMessage) + "\nActual: " + fmt.Sprint(mt))
	}

	if string(m) != expected {
		t.Error("Expected: " + string(expected) + "\nActual: " + string(m))
	}

	connection.WriteControl(websocket.CloseMessage, []byte{}, time.Now().Add(10*time.Second))
}
