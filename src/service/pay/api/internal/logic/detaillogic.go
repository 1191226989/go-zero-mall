package logic

import (
	"context"

	"mall/service/pay/api/internal/svc"
	"mall/service/pay/api/internal/types"
	"mall/service/pay/rpc/pay"

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
	ret, err := l.svcCtx.PayRpc.Detail(l.ctx, &pay.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DetailResponse{
		Id:     ret.Id,
		Uid:    ret.Uid,
		Oid:    ret.Oid,
		Amount: ret.Amount,
		Source: ret.Source,
		Status: ret.Status,
	}, nil
}
