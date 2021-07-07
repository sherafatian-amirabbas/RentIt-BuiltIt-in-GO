package websocket

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/cs-ut-ee/hw2-group-3/pkg/service"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	service service.IService
	router  *mux.Router
}

type Message struct {
	Resource string            `json:"resource"`
	Params   map[string]string `json:"params"`
}

func NewWebsocketHandler(service service.IService, router *mux.Router) *WebsocketHandler {
	return &WebsocketHandler{
		service: service,
		router:  router,
	}
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (handler *WebsocketHandler) handshake(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if websocket.IsUnexpectedCloseError(err, websocket.CloseNoStatusReceived) {
			log.Println("read:", err)
			break
		}

		result, handleErr := handler.handleMessage(message)
		if handleErr != nil {
			log.Println("Message handling error:", err)
		}

		if handleErr == nil {
			err = c.WriteMessage(mt, result)
		} else {
			err = c.WriteMessage(mt, []byte(handleErr.Error()))
		}

		if websocket.IsUnexpectedCloseError(err, websocket.CloseNoStatusReceived) {
			log.Println("write:", err)
			break
		}
	}
}

func (handler *WebsocketHandler) RegisterRoutes() {
	handler.router.HandleFunc("/websocket", handler.handshake).Methods(http.MethodGet)
}

func (handler *WebsocketHandler) handleMessage(message []byte) ([]byte, error) {
	var request Message
	err := json.Unmarshal(message, &request)
	if err != nil {
		return nil, handleError(err, "Could not read message: ")
	}

	switch request.Resource {
	case "plants":
		result, err := handler.service.GetAllPlants()
		if err != nil {
			return nil, handleError(err, "Service error: ")
		}

		return json.Marshal(result)
	case "plant/price":
		id, err := strconv.ParseInt(request.Params["plantId"], 10, 64)
		if err != nil {
			return nil, handleError(err, "Invalid parameter 'plantId': ")
		}

		from, err := time.Parse("2006-01-02", request.Params["from"])
		if err != nil {
			return nil, handleError(err, "Invalid parameter 'from': ")
		}

		to, err := time.Parse("2006-01-02", request.Params["to"])
		if err != nil {
			return nil, handleError(err, "Invalid parameter 'to': ")
		}

		result, err := handler.service.GetPlantPrice(id, from, to)
		if err != nil {
			return nil, handleError(err, "Service error: ")
		}

		return json.Marshal(result)
	case "plant/available":
		id, err := strconv.ParseInt(request.Params["plantId"], 10, 64)
		if err != nil {
			return nil, handleError(err, "Invalid parameter 'plantId': ")
		}

		from, err := time.Parse("2006-01-02", request.Params["from"])
		if err != nil {
			return nil, handleError(err, "Invalid parameter 'from': ")
		}

		to, err := time.Parse("2006-01-02", request.Params["to"])
		if err != nil {
			return nil, handleError(err, "Invalid parameter 'to': ")
		}

		result, err := handler.service.IsPlantAvailable(id, from, to)
		if err != nil {
			return nil, handleError(err, "Service error: ")
		}

		return json.Marshal(result)

	default:
		return []byte("Ooopsy doopsie. You made a mess"), nil
	}
}

func handleError(err error, prefix string) error {
	if err != nil {
		return errors.New(prefix + err.Error())
	}

	return err
}
