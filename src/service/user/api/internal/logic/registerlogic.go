package logic

import (
	"context"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"
	"mall/service/user/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	ret, err := l.svcCtx.UserRpc.Register(l.ctx, &proto3.RegisterRequest{
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		Id:     ret.Id,
		Name:   ret.Name,
		Gender: ret.Gender,
		Mobile: ret.Mobile,
	}, nil
}
