package logic

import (
	"context"

	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *proto3.UpdateRequest) (*proto3.UpdateResponse, error) {
	// 查询订单是否存在
	ret, err := l.svcCtx.OrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Uid != 0 {
		ret.Uid = in.Uid
	}
	if in.Pid != 0 {
		ret.Pid = in.Pid
	}
	if in.Amount != 0 {
		ret.Amount = in.Amount
	}
	if in.Status != 0 {
		ret.Status = in.Status
	}

	err = l.svcCtx.OrderModel.Update(l.ctx, ret)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &proto3.UpdateResponse{}, nil
}
