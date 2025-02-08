package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrades = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handelConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrades.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket 업그레이드 실패", err)
	}
	defer conn.Close()

}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8080", nil)
}
