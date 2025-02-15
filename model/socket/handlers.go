package model

import (
	"fmt"
	"net/http"
	"sync"

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
	//함수가 끝날떄 자동 실행
	defer conn.Close()

}

// 클라이언트 요청 구조체
type Client struct {
	conn *websocket.Conn
	send chan []byte
}

type ClientHub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}
