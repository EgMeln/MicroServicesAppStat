package user

import (
	"github.com/EgMeln/game_system/user_service/app/iternal/postgreSQL"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetAllUser(c echo.Context) error {
	return c.String(http.StatusOK, postgreSQL.GetAllUser())
}
func GetUserById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, postgreSQL.GetUserByID(id))
}
func AddUser(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	postgreSQL.AddUser(username, email, password)
	return c.String(http.StatusOK, postgreSQL.AddUser(username, email, password))
}
func DeleteUser(c echo.Context) error {
	username := c.FormValue("username")
	return c.String(http.StatusOK, postgreSQL.DeleteUser(username))
}
