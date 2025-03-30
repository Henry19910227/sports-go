package user

import (
	"api_gateway/engine"
	"api_gateway/factory/controller"
)

func SetRoute(group *engine.RouterGroup, factory controller.Factory) {
	userController := factory.UserController()
	group.EndPoint("1001", userController.Unmarshal, userController.Login)
}
