package middleware

import (
	"api_gateway/engine"
)

type Controller interface {
	UnMarshalData(ctx *engine.Context)
	//Transaction(db *gorm.DB) game.HandlerFunc
	Function1(ctx *engine.Context)
	Function2(ctx *engine.Context)
}
