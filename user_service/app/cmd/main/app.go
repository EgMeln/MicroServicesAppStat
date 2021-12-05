package main

import (
	"fmt"
	"github.com/EgMeln/game_system/user_service/app/iternal/GHandlers"
	"github.com/EgMeln/game_system/user_service/app/iternal/grpcProto"
	"github.com/EgMeln/game_system/user_service/app/iternal/services"
	"github.com/EgMeln/game_system/user_service/app/iternal/user"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "sing in or sign up please")
	})
	e.GET("/allUsers", user.GetAllUser)
	e.GET("/users/:id", user.GetUserById)
	/*http://localhost:1321/addUser?username=admin&email=deee&password=fefe222*/
	e.GET("/addUser", user.AddUser)
	e.GET("/userDelete", user.DeleteUser)

	fmt.Println("add rest or grpc")
	var str string

	fmt.Scan(&str)

	switch str {
	case "rest":
		e.Logger.Fatal(e.Start(":1321"))

	case "grpc":
		server := grpc.NewServer()
		handler := &GHandlers.Handler{}
		address, err := net.Listen("tcp", ":1321")
		if err != nil {
			fmt.Println("ERROR LISTEN")
		}
		reflection.Register(server)
		s := &services.Server{}
		services.GameConnect(s)
		services.InfoConnect(s)

		grpcProto.RegisterUserServiceServer(server, handler)
		err = server.Serve(address)
		if err := server.Serve(address); err != nil {
			fmt.Println("ERROR SERVE")
		}

	}
}
