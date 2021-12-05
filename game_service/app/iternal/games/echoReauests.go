package games

import (
	"github.com/EgMeln/game_system/game_service/app/iternal/db"
	"github.com/EgMeln/game_system/game_service/app/iternal/gameSystem"
	"github.com/labstack/echo"
	"net/http"
)

func Game32(c echo.Context) error {
	str, winHeroes := gameSystem.Run32()
	return c.String(http.StatusOK, db.PostgreSQL(str, *winHeroes))
}
func Game64(c echo.Context) error {
	str, winHeroes := gameSystem.Run64()
	return c.String(http.StatusOK, db.PostgreSQL(str, *winHeroes))
}
func Game128(c echo.Context) error {
	str, winHeroes := gameSystem.Run128()
	return c.String(http.StatusOK, db.PostgreSQL(str, *winHeroes))
}
