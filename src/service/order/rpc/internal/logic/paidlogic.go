package logic

import (
	"context"

	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *proto3.PaidRequest) (*proto3.PaidResponse, error) {
	// 查询订单是否存在
	ret, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	// 修改订单状态
	ret.Status = 1
	err = l.svcCtx.OrderModel.Update(l.ctx, ret)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &proto3.PaidResponse{}, nil
}
