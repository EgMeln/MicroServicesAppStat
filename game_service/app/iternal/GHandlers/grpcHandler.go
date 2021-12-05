package GHandlers

import (
	"context"
	"github.com/EgMeln/game_system/game_service/app/iternal/db"
	"github.com/EgMeln/game_system/game_service/app/iternal/gameSystem"
	"github.com/EgMeln/game_system/game_service/app/iternal/grpcProto"
)

type Handler struct {
	grpcProto.UnimplementedGameServiceServer
}

func (g *Handler) Game32(context.Context, *grpcProto.GameEmpty) (*grpcProto.LogFight, error) {
	str, winHeroes := gameSystem.Run32()
	return &grpcProto.LogFight{FullFight: db.PostgreSQL(str, *winHeroes)}, nil
}
func (g *Handler) Game64(context.Context, *grpcProto.GameEmpty) (*grpcProto.LogFight, error) {
	str, winHeroes := gameSystem.Run64()
	return &grpcProto.LogFight{FullFight: db.PostgreSQL(str, *winHeroes)}, nil
}
func (g *Handler) Game128(context.Context, *grpcProto.GameEmpty) (*grpcProto.LogFight, error) {
	str, winHeroes := gameSystem.Run128()
	return &grpcProto.LogFight{FullFight: db.PostgreSQL(str, *winHeroes)}, nil
}
