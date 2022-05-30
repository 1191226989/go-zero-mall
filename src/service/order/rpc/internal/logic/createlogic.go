package logic

import (
	"context"
	"database/sql"
	"fmt"

	"mall/service/order/model"
	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/proto3"
	"mall/service/user/rpc/user"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

func (l *CreateLogic) Create(in *proto3.CreateRequest) (*proto3.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, status.Error(100, "用户不存在")
	}
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	if err := barrier.CallWithDB(db, func(tx *sql.Tx) error {
		newOrder := model.Order{
			Uid:     in.Uid,
			Pid:     in.Pid,
			Amount:  in.Amount,
			Status:  0,
			OrderNo: in.OrderNo,
		}
		// 创建订单
		_, err = l.svcCtx.OrderModel.TxInsert(tx, &newOrder)
		if err != nil {
			return fmt.Errorf("order.rpc 订单创建失败:" + err.Error())
		}

		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &proto3.CreateResponse{}, nil
}
