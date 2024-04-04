package factories

import (
	"go-url-shortener/src/controllers"
	"go-url-shortener/src/services"
)

var userService = services.NewUserService()
var UserController = controllers.NewUserController(*userService)
