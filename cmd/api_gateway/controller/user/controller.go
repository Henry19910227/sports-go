package user

import (
	"api_gateway/client/user"
	"api_gateway/engine"
	"fmt"
	"google.golang.org/protobuf/proto"
	"sports-go/shared/pb/api_gateway"
)

type controller struct {
	userClient user.Client
}

func New(client user.Client) Controller {
	return &controller{client}
}

func (c *controller) Unmarshal(ctx *engine.Context) {
	mid := ctx.MustGet("mid").(uint16)
	sid := ctx.MustGet("sid").(uint16)
	payload := ctx.MustGet("payload").([]byte)
	var pb proto.Message
	switch {
	case mid == 500 && sid == 1001:
		pb = &api_gateway.LoginReq{}
	default:
		fmt.Println("No matching case for the given mid and sid")
		ctx.Abort()
		return
	}
	if err := proto.Unmarshal(payload, pb); err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("pb", pb)
}

func (c *controller) Login(ctx *engine.Context) {
	login := ctx.MustGet("pb").(*api_gateway.LoginReq)
	output, err := c.userClient.Login(ctx.Ctx(), login.Uid, login.Token)
	if err != nil {
		fmt.Println("Error:", err)
	}
	data, _ := ctx.MarshalData(500, 1001, output)
	ctx.WriteData(data)
}
