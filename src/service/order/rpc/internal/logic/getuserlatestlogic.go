package logic

import (
	"context"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type GetUserLatestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLatestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLatestLogic {
	return &GetUserLatestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLatestLogic) GetUserLatest(in *proto3.GetUserLatestRequest) (*proto3.GetUserLatestResponse, error) {
	ret, err := l.svcCtx.OrderModel.FindOneByUid(in.Uid)
	if err != nil {
		return nil, status.Error(100, "订单不存在")
	}
	return &proto3.GetUserLatestResponse{
		Id:     ret.Id,
		Uid:    ret.Uid,
		Pid:    ret.Pid,
		Amount: ret.Amount,
		Status: ret.Status,
	}, nil
}
