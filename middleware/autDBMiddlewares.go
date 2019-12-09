package middlewares

import (
	"github.com/labstack/echo"
	"go-training-restful/config"
	"go-training-restful/models"
)

var db = config.DB

func BasicAuthDB (username, password string, c echo.Contect)(bool, error){
var db = config.DB
var user models.user
if err := db.Where

}