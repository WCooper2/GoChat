/*
`* File: backend/main.go
 * Basic service of realtime chat between a UI Client and this web server.
 * Opens and listens on a websocket for properly routed connections
 * Will Cooper
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var websocketController = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	//This will allow all requests. Future will be to add authentication and accounts
	CheckOrigin: func(r *http.Request) bool { return true },
}

func buildRoutes() {
	http.HandleFunc("/", basicHandler)
	http.HandleFunc("/ws", supplyWS)
}

func socketReader(conn *websocket.Conn) {
	for {
		//Read a message
		messageType, contents, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		//Display the message
		fmt.Println(string(contents))

		//send the message back as a simple verification
		sendErr := conn.WriteMessage(messageType, contents)
		if sendErr != nil {
			log.Println(sendErr)
			return
		}
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connection Successful")
}

func supplyWS(w http.ResponseWriter, r *http.Request) {
	socket, err := websocketController.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// Read indefinetly
	socketReader(socket)
}

func main() {
	fmt.Println("Connecting to webchat")
	buildRoutes()
	http.ListenAndServe(":8080", nil)
}
