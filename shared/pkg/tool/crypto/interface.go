package crypto

type Tool interface {
	Marshal(mid uint16, sid uint16, data []byte) ([]byte, error)
	UnMarshal(data []byte) (uint16, uint16, []byte, error)
}
