package logic

import (
	"context"
	"database/sql"
	"fmt"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/proto3"
	"mall/service/user/rpc/user"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRevertLogic) CreateRevert(in *proto3.CreateRequest) (*proto3.CreateResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, status.Error(100, "用户不存在")
	}
	// 获取存储dtm_barrier的mysql数据库
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
		// 通过orderNo获取创建的订单
		resOrder, err := l.svcCtx.OrderModel.FindOneByOrderNo(in.OrderNo)
		if err != nil {
			return fmt.Errorf("订单不存在")
		}
		// 修改订单状态9，标识订单已失效，并更新订单
		resOrder.Status = 9
		err = l.svcCtx.OrderModel.TxUpdate(tx, resOrder)
		if err != nil {
			return fmt.Errorf("order.rpc 订单更新失败:" + err.Error())
		}

		return nil
	}); err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &proto3.CreateResponse{}, nil
}
