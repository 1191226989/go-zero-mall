package logic

import (
	"context"

	"mall/service/product/api/internal/svc"
	"mall/service/product/api/internal/types"
	"mall/service/product/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.RemoveRequest) (resp *types.RemoveResponse, err error) {
	_, err = l.svcCtx.ProductRpc.Remove(l.ctx, &proto3.RemoveRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &types.RemoveResponse{}, nil
}
