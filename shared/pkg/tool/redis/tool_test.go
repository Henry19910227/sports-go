package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestTool_SetEX(t *testing.T) {
	redisTool := New()
	if err := redisTool.SetEX("Henry", "Hello!", time.Second*10); err != nil {
		fmt.Println(err)
	}
}

func TestTool_LPush(t *testing.T) {
	redisTool := New()
	if err := redisTool.LPush("round_infos", "1", "2", "3"); err != nil {
		fmt.Println(err)
	}
}

func TestTool_RPush(t *testing.T) {
	redisTool := New()
	if err := redisTool.RPush("round_infos", "1", "2", "3"); err != nil {
		fmt.Println(err)
	}
}

func TestTool_LRange(t *testing.T) {
	redisTool := New()
	val := redisTool.LRange("round_infos", 0, -1)
	fmt.Println(val)
}
