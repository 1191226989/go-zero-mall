package logic

import (
	"context"

	"mall/service/product/model"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/proto3"

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
	// 查询产品是否存在
	ret, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	if in.Name != "" {
		ret.Name = in.Name
	}
	if in.Desc != "" {
		ret.Desc = in.Desc
	}
	if in.Stock != 0 {
		ret.Stock = in.Stock
	}
	if in.Amount != 0 {
		ret.Amount = in.Amount
	}
	if in.Status != 0 {
		ret.Status = in.Status
	}

	err = l.svcCtx.ProductModel.Update(l.ctx, ret)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &proto3.UpdateResponse{}, nil
}
