package crypto_json

import "google.golang.org/protobuf/proto"

type Tool interface {
	Marshal(mid int, sid int, payload proto.Message) ([]byte, error)
	Mid(b []byte) int
	Sid(b []byte) int
	Payload(b []byte, out proto.Message) error
}
