package util

import (
	"encoding/json"
	"strconv"
	"strings"
)

func Parser(input interface{}, output interface{}) error {
	marshal, err := json.Marshal(input)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(marshal, output); err != nil {
		return err
	}
	return nil
}

func StringToInt32Array(input string, sep string) []int32 {
	elements := make([]int32, 0)
	for _, s := range strings.Split(input, sep) {
		i, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		elements = append(elements, int32(i))
	}
	return elements
}
