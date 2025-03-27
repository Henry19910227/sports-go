package crypto_json

import (
	"encoding/base64"
	"encoding/json"
	"game-go/core/model/message"
	"google.golang.org/protobuf/proto"
)

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) Marshal(mid int, sid int, payload proto.Message) ([]byte, error) {

	// 轉成 pb data
	pbData, err := proto.Marshal(payload)
	if err != nil {
		return nil, err
	}

	// 設置 payload
	msg := message.Model{}
	msg.Mid = mid
	msg.Sid = sid
	msg.Payload = base64.StdEncoding.EncodeToString(pbData)

	// 轉成 json data
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (t *tool) Mid(b []byte) int {
	var m message.Model
	if err := json.Unmarshal(b, &m); err != nil {
		return 0
	}
	return m.Mid
}

func (t *tool) Sid(b []byte) int {
	var m message.Model
	if err := json.Unmarshal(b, &m); err != nil {
		return 0
	}
	return m.Sid
}

func (t *tool) Payload(b []byte, out proto.Message) error {
	var m message.Model
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	data, err := base64.StdEncoding.DecodeString(m.Payload)
	if err != nil {
		return err
	}
	if err := proto.Unmarshal(data, out); err != nil {
		return err
	}
	return nil
}
