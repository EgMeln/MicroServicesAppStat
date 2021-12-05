package info

import (
	"github.com/EgMeln/game_system/info_service/app/iternal/db"
	"github.com/labstack/echo"
	"net/http"
)

func StrGame32(c echo.Context) error {
	return c.String(http.StatusOK, db.Sql(1))
}
func StrGame64(c echo.Context) error {
	return c.String(http.StatusOK, db.Sql(2))
}
func StrGame128(c echo.Context) error {
	return c.String(http.StatusOK, db.Sql(3))
}
func AllUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Command to print all users")
}
func UserByID(c echo.Context) error {
	return c.String(http.StatusOK, "Command to print user by id")
}
func AddUser(c echo.Context) error {
	return c.String(http.StatusOK, "Command to add user")
}
func UserDelete(c echo.Context) error {
	return c.String(http.StatusOK, "Command to delete user")
}
