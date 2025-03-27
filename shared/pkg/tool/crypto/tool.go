package crypto

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) Marshal(mid uint16, sid uint16, data []byte) ([]byte, error) {
	// 創建 buffer 以存儲頭部和 proto 數據
	buf := new(bytes.Buffer)

	// 將 mid 和 sid 寫入 buffer
	if err := binary.Write(buf, binary.BigEndian, mid); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, sid); err != nil {
		return nil, err
	}
	// 寫入 proto 數據
	if _, err := buf.Write(data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *tool) UnMarshal(data []byte) (uint16, uint16, []byte, error) {
	if len(data) < 4 {
		return 0, 0, nil, fmt.Errorf("data too short")
	}

	// 讀取 mid 和 sid
	buf := bytes.NewReader(data)
	var mid, sid uint16
	if err := binary.Read(buf, binary.BigEndian, &mid); err != nil {
		return 0, 0, nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &sid); err != nil {
		return 0, 0, nil, err
	}

	// 提取剩餘的 Proto 數據
	message := data[4:]

	return mid, sid, message, nil
}
