package crypto_json

import (
	"fmt"
	"game-go/shared/res"
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	pb := &res.UserData{}
	pb.UserId = 1
	pb.Score = 10000

	// 轉成 crypto
	msg, err := New().Marshal(1, 2, pb)
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}

	fmt.Println(string(msg))

	// 解出 mid
	fmt.Println(New().Mid(msg))

	// 解出 sid
	fmt.Println(New().Sid(msg))

	// 解出 payload
	var u = res.UserData{}
	if err := New().Payload(msg, &u); err != nil {
		log.Fatalf("Failed to Payload: %v", err)
	}
	fmt.Println(u.UserId)
	fmt.Println(u.Score)
}
