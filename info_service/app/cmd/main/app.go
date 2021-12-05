package main

import (
	"fmt"
	"github.com/EgMeln/game_system/info_service/app/iternal/GHandlers"
	pb1 "github.com/EgMeln/game_system/info_service/app/iternal/grpcProto"
	"github.com/EgMeln/game_system/info_service/app/iternal/info"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Here you can play a simple game.\nWrite /game32,/game64,/game128 in order to start collecting statistics")
	})
	e.GET("/info/game32", info.StrGame32)
	e.GET("/info/game64", info.StrGame64)
	e.GET("/info/game128", info.StrGame128)
	e.GET("/allUsers", info.AllUsers)
	e.GET("/users/:id", info.UserByID)
	e.GET("/addUser", info.AddUser)
	e.GET("/userDelete,", info.UserDelete)

	fmt.Println("add rest or grpcProto")
	var str string

	fmt.Scan(&str)

	switch str {
	case "rest":
		e.Logger.Fatal(e.Start(":1323"))
	case "grpc":
		server := grpc.NewServer()
		handler := &GHandlers.Handler{}
		pb1.RegisterInfoServiceServer(server, handler)
		address, err := net.Listen("tcp", ":1323")
		if err != nil {
			fmt.Println("ERROR LISTEN")
		}
		reflection.Register(server)
		if err := server.Serve(address); err != nil {
			fmt.Println("ERROR SERVE")
		}

	}
}
