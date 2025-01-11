package main

import (
	"context"
	"log"
	"net/http"

	"github.com/duolok/wisp/models"
)

func (m MLWrapper) executeModel(w http.ResponseWriter, r *http.Request) {
	modelName := r.URL.Query().Get("model")
	streaming := StringToBool(r.URL.Query().Get("streaming"))

	conn, err := websocketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading messaagae: ", err)
			return
		}

		if streaming {
			response, err := m.wrapper.LoadStreamingModel(modelName, string(msg))

			processFunc := func(ctx context.Context, part []byte) error {
				err = conn.WriteMessage(msgType, part)
				if err != nil {
					log.Println("Error writing to websocket: ", err)
					return err
				}

				return nil
			}

			err = models.CallStreamingOutputFunction(modelName, response, processFunc)
			if err != nil {
				log.Fatal("streaming output processing error: ", err)
			}

		} else {
			modelResponse, err := m.wrapper.LoadModel(modelName, string(msg))
			err = conn.WriteMessage(msgType, []byte(modelResponse))
			if err != nil {
				log.Println("Error writing to websocket: ", err)
				return
			}
		}
	}
}
