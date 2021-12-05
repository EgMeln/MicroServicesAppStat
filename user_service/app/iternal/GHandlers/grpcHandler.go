package GHandlers

import (
	"context"
	"github.com/EgMeln/game_system/user_service/app/iternal/postgreSQL"
	pb3 "github.com/EgMeln/game_system/user_service/app/iternal/services/game/grpcProto"
	pb2 "github.com/EgMeln/game_system/user_service/app/iternal/services/info/grpcProto"
	pb1 "github.com/EgMeln/game_system/user_service/app/iternal/services/user/grpcProto"

	"github.com/labstack/echo/v4"
	"strconv"
)

type Handler struct {
	*echo.Echo
	pb1.UnimplementedUserServiceServer
	pb2.InfoServiceClient
	pb3.GameServiceClient
}

func (g *Handler) GetUserById(ctx context.Context, id *pb1.UserId) (*pb1.LogUser, error) {
	return &pb1.LogUser{ResultStr: postgreSQL.GetUserByID(strconv.Itoa(int(id.Id)))}, nil
}
func (g *Handler) GetAllUser(ctx context.Context, in *pb1.Empty) (*pb1.LogUser, error) {
	return &pb1.LogUser{ResultStr: postgreSQL.GetAllUser()}, nil
}
func (g *Handler) AddUser(ctx context.Context, in *pb1.UserAdd) (*pb1.LogUser, error) {
	return &pb1.LogUser{ResultStr: postgreSQL.AddUser(in.Username, in.Email, in.Password)}, nil
}
func (g *Handler) DeleteUser(ctx context.Context, in *pb1.UserDelete) (*pb1.LogUser, error) {
	return &pb1.LogUser{ResultStr: postgreSQL.DeleteUser(in.Username)}, nil
}
