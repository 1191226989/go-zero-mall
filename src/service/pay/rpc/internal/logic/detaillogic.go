package logic

import (
	"context"

	"mall/service/pay/model"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/proto3"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *proto3.DetailRequest) (*proto3.DetailResponse, error) {
	// 查询支付是否存在
	ret, err := l.svcCtx.PayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "支付记录不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	return &proto3.DetailResponse{
		Id:     ret.Id,
		Uid:    ret.Uid,
		Oid:    ret.Oid,
		Amount: ret.Amount,
		Source: ret.Source,
		Status: ret.Status,
	}, nil
}
