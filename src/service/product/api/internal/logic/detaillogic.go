package logic

import (
	"context"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"
	"mall/service/product/rpc/proto3"

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
	ret, err := l.svcCtx.ProductRpc.Detail(l.ctx, &proto3.DetailRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DetailResponse{
		Id:     ret.Id,
		Name:   ret.Name,
		Desc:   ret.Desc,
		Stock:  ret.Stock,
		Amount: ret.Amount,
		Status: ret.Status,
	}, nil
}
