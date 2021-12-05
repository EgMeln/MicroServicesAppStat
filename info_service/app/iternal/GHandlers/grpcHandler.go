package GHandlers

import (
	"context"
	"github.com/EgMeln/game_system/info_service/app/iternal/db"
	pb1 "github.com/EgMeln/game_system/info_service/app/iternal/grpcProto"
)

type Handler struct {
	pb1.UnimplementedInfoServiceServer
}

func (g *Handler) StrGame32(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(1)}, nil
}
func (g *Handler) StrGame64(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(2)}, nil
}
func (g *Handler) StrGame128(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(3)}, nil
}
func (g *Handler) AllUsers(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(4)}, nil
}
func (g *Handler) UserByID(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(5)}, nil
}
func (g *Handler) AddUser(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(6)}, nil
}
func (g *Handler) UserDelete(ctx context.Context, in *pb1.InfoEmpty) (*pb1.AnswerInfo, error) {
	return &pb1.AnswerInfo{AnsSTR: db.Sql(7)}, nil

}
