package services

import (
	"fmt"
	gameService "github.com/EgMeln/game_system/user_service/app/iternal/services/game/grpcProto"
	infoService "github.com/EgMeln/game_system/user_service/app/iternal/services/info/grpcProto"
	"google.golang.org/grpc"
)

type Server struct {
	GameService gameService.GameServiceClient
	InfoService infoService.InfoServiceClient
}

func GameConnect(s *Server) {
	cnnUser, err := grpc.Dial("localhost:1322", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("[IN] %s.GAmeConnect: ошибка подключения к сервису GAme", err)
	}

	s.GameService = gameService.NewGameServiceClient(cnnUser)
}
func InfoConnect(s *Server) {
	cnnUser, err := grpc.Dial("localhost:1323", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("[IN] %s.InfoConnect: ошибка подключения к сервису Info", err)
	}

	s.InfoService = infoService.NewInfoServiceClient(cnnUser)
}
