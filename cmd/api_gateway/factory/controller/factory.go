package controller

import (
	middController "api_gateway/controller/middleware"
	userController "api_gateway/controller/user"
	"api_gateway/factory/client"
)

type factory struct {
	clientFactory client.Factory
}

func New(clientFactory client.Factory) Factory {
	controllerFactory := &factory{
		clientFactory: clientFactory,
	}
	return controllerFactory
}

func (f *factory) UserController() userController.Controller {
	return userController.New(f.clientFactory.UserClient())
}

func (f *factory) MiddController() middController.Controller {
	return middController.New()
}
