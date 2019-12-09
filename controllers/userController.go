package controllers
import (
	"net/http"
	"go-training-restful/lib/database"
	"github.com/labstack/echo"

)

//GetUserController ... 
func GetUsersController(e echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError (http.StatusBadRequest,
		err.Error())
	}
	return e.JSON(http.StatusOK, users)
}