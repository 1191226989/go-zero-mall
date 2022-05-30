package logic

import (
	"context"
	"encoding/json"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"
	"mall/service/user/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// 通过 l.ctx.Value("uid") 可获取 jwt token 中自定义的参数
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	ret, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &proto3.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.UserInfoResponse{
		Id:     ret.Id,
		Name:   ret.Name,
		Gender: ret.Gender,
		Mobile: ret.Mobile,
	}, nil
}
