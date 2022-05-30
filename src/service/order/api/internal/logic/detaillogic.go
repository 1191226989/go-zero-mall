package logic

import (
	"context"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	ret, err := l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     ret.Id,
		Uid:    ret.Uid,
		Pid:    ret.Pid,
		Amount: ret.Amount,
		Status: ret.Status,
	}, nil
}
