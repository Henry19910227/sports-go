package controller

import (
	middController "api_gateway/controller/middleware"
	userController "api_gateway/controller/user"
)

type Factory interface {
	UserController() userController.Controller
	MiddController() middController.Controller
}
