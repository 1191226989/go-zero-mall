package logic

import (
	"context"

	"mall/service/product/model"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建产品
func (l *CreateLogic) Create(in *proto3.CreateRequest) (*proto3.CreateResponse, error) {
	newPrduct := model.Product{
		Name:   in.Name,
		Desc:   in.Desc,
		Stock:  in.Stock,
		Amount: in.Amount,
		Status: in.Status,
	}
	ret, err := l.svcCtx.ProductModel.Insert(l.ctx, &newPrduct)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newPrduct.Id, err = ret.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &proto3.CreateResponse{
		Id: newPrduct.Id,
	}, nil
}
