package logic

import (
	"context"

	"mall/common/randx"
	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"
	"mall/service/order/rpc/order"
	"mall/service/product/rpc/product"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 提示：SagaGrpc.Add 方法第一个参数 action 是微服务 grpc 访问的方法路径，这个方法路径需要分别去以下文件中寻找
// mall/service/order/rpc/order/order.pb.go
// mall/service/product/rpc/product/product.pb.go
// 按关键字 Invoke 搜索即可找到。
func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// 获取 OrderRpc BuildTarget
	orderRpcBusiServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// 获取 ProductRpc BuildTarget
	productRpcBusiServer, err := l.svcCtx.Config.ProductRpc.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "订单创建异常")
	}

	// 生成订单号
	orderNo := randx.MustGenOrderNo(req.Uid)
	// dtm 服务的 etcd 注册地址
	var dtmServer = l.svcCtx.Config.DtmServer.Address
	// 创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	// 创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcBusiServer+"/order.Order/Create", orderRpcBusiServer+"/order.Order/CreateRevert", &order.CreateRequest{
			Uid:     req.Uid,
			Pid:     req.Pid,
			Amount:  req.Amount,
			OrderNo: orderNo,
			Status:  0,
		}).
		Add(productRpcBusiServer+"/product.Product/DecrStock", productRpcBusiServer+"/product.Product/DecrStockRevert", &product.DecrStockRequest{
			Id:  req.Pid,
			Num: 1,
		})

	// 事务提交
	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{
		OrderNo: orderNo,
	}, nil
}
