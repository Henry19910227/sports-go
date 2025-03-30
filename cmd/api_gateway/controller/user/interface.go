package user

import "api_gateway/engine"

type Controller interface {
	Unmarshal(ctx *engine.Context)
	Login(ctx *engine.Context)
}
