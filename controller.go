package main

import (
    "log"
    "net/http"
)

func (m MLWrapper) executeModel (w http.ResponseWriter, r *http.Request) {
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
