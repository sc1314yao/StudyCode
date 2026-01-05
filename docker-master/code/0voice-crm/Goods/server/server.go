package server

import (
	"Goods/config"
	"Goods/proto"
	"context"
)

type goodsServer struct {
	config *config.Config
	proto.UnimplementedGoodsServer
}

func NewGoodsServer(config *config.Config) proto.GoodsServer {
	return &goodsServer{
		config: config,
	}
}

func (s *goodsServer) Get(ctx context.Context, in *proto.GetGoodsReq) (*proto.GetGoodsRes, error) {
	return &proto.GetGoodsRes{
		GoodsId:       in.GoodsId,
		GoodsName:     s.config.Goods.Name,
		GoodsCategory: s.config.Goods.Category,
	}, nil
}
