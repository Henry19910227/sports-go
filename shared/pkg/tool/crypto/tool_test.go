package crypto

import (
	"fmt"
	"game-go/shared/req"
	"game-go/shared/res"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	user := &res.ErrorMessage{}
	user.Code = 400

	pb, err := proto.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 轉成 crypto
	data, err := New().Marshal(7, 107, pb)
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}
	//Binary(data)
	HexBinary(data)

	// 解出 payload
	var u = res.ErrorMessage{}
	mid, sid, payload, err := New().UnMarshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = proto.Unmarshal(payload, &u)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mid)
	fmt.Println(sid)
	fmt.Println(u.Code)
	fmt.Println(u.Desc)
}

func TestLoginMarshalToBinary(t *testing.T) {
	login := &req.LoginReq{}
	login.AgentName = "wali-core"
	login.Token = "9:Cv4XkLoA"
	login.Platform = 6
	login.RequestId = 0
	login.Nickname = "Henry"
	login.Version = "1"

	pb, err := proto.Marshal(login)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 轉成 crypto
	data, err := New().Marshal(7, 7, pb)
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}
	HexBinary(data)
}

func TestMarshalToBinary(t *testing.T) {
	data, err := New().Marshal(99, 100, []byte{})
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}
	HexBinary(data)
}

func Binary(data []byte) {
	// 1 byte 可以表示 8 個 2 進位數字
	for _, b := range data {
		fmt.Printf("%08b ", b)
	}
}

func HexBinary(data []byte) {
	// 一個 16 進位數字需用 4 bit 表示，所以 1 byte 可以表示 2 個 16 進位數字
	// 0x00 0x01 0x27 x010 = 0(int8), 1(int8), 39(int8), 16(int8),
	// 0x0001 0x2710 = 1(int16) 10000(int16)
	for _, b := range data {
		fmt.Printf("%02x ", b)
	}
}
