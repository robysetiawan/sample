package routes

import (	
     "go-training-restful/controllers"
	"github.com/labstack/echo"
)
//New ...
func New() *echo.Echo{
	e:= echo.New()
	 e.GET ("/users", controllers.GetUsersController)
	return e	
}