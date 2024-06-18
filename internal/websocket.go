package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Websocket struct {
	upgrader    websocket.Upgrader
	infoService *Info
}

func NewWebsocket(infoService *Info) *Websocket {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &Websocket{upgrader: upgrader, infoService: infoService}
}

func (webs *Websocket) HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := webs.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	fmt.Println("Client connected")
	ticker := time.NewTicker(1 * time.Second)

	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			msg := fmt.Sprintf("Current time: %s", t.Format(time.RFC3339))
			fmt.Println(msg)
			if err := ws.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				fmt.Println("Error writing message: ", err)
				return
			}
		}
	}
}
