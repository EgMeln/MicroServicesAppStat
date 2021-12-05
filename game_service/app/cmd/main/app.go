package main

import (
	"fmt"
	"github.com/EgMeln/game_system/game_service/app/iternal/GHandlers"
	"github.com/EgMeln/game_system/game_service/app/iternal/games"
	"github.com/EgMeln/game_system/game_service/app/iternal/grpcProto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Are you ready to collect statistics?\nPlease enter the command")
	})
	e.GET("/game32", games.Game32)
	e.GET("/game64", games.Game64)
	e.GET("/game128", games.Game128)

	fmt.Println("add rest or grpc")
	var str string

	fmt.Scan(&str)

	switch str {
	case "rest":
		e.Logger.Fatal(e.Start(":1322"))
	case "grpc":
		server := grpc.NewServer()
		handler := &GHandlers.Handler{}
		grpcProto.RegisterGameServiceServer(server, handler)
		address, err := net.Listen("tcp", ":1322")
		if err != nil {
			fmt.Println("ERROR LISTEN")
		}
		reflection.Register(server)
		if err := server.Serve(address); err != nil {
			fmt.Println("ERROR SERVE")
		}
	}

}
