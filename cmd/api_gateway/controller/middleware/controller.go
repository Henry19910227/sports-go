package middleware

import (
	"api_gateway/engine"
	"fmt"
	"sports-go/shared/pkg/tool/crypto"
)

type controller struct {
}

func New() Controller {
	return &controller{}
}

func (c *controller) UnMarshalData(ctx *engine.Context) {
	mid, sid, payload, err := crypto.New().UnMarshal(ctx.RawData())
	if err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("mid", mid)
	ctx.Set("sid", sid)
	ctx.Set("payload", payload)
}

//func (c *controller) Transaction(db *gorm.DB) game.HandlerFunc {
//	return func(ctx *game.Context) {
//		txHandle := db.Begin()
//		defer func() {
//			txHandle.Rollback()
//		}()
//		ctx.Set("tx", txHandle)
//		ctx.Next()
//	}
//}

func (c *controller) Function1(ctx *engine.Context) {
	defer func() {
		fmt.Println("Function1 end")
	}()
	fmt.Println("Function1 begin")
	ctx.Next()
}

func (c *controller) Function2(ctx *engine.Context) {
	defer func() {
		fmt.Println("Function2 end")
	}()
	fmt.Println("Function2 begin")
	ctx.Next()
}
