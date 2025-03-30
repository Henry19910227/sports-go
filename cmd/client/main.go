package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/url"
	"sports-go/shared/pb/api_gateway"
	"sports-go/shared/pkg/tool/crypto"
)

func main() {
	// 建立 WebSocket 連線
	serverURL := url.URL{Scheme: "ws", Host: "localhost:7070", Path: "/game"}
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		log.Fatalf("無法連線到伺服器: %v", err)
	}
	defer conn.Close()

	// 接收訊息
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("讀取訊息錯誤:", err)
				return
			}
			log.Printf("收到訊息: %s", message)
		}
	}()

	req := &api_gateway.LoginReq{
		Uid:   10001,
		Token: "123456",
	}
	pb, err := proto.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	result, err := crypto.New().Marshal(500, 1001, pb)
	if err != nil {
		fmt.Println(err)
	}
	// 傳送訊息
	if err := conn.WriteMessage(websocket.BinaryMessage, result); err != nil {
		fmt.Println(err)
	}
}
