package logic

import (
	"context"

	"mall/service/order/rpc/order"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/proto3"
	"mall/service/user/model"
	"mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *proto3.CallbackRequest) (*proto3.CallbackResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 查询支付记录是否存在
	ret, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "支付记录不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	// 支付金额与订单金额不符
	if in.Amount != ret.Amount {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}

	// 更新支付记录数据
	ret.Source = in.Source
	ret.Status = in.Status
	err = l.svcCtx.PayModel.Update(l.ctx, ret)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 更新订单支付状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{
		Id: in.Oid,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &proto3.CallbackResponse{}, nil
}
