package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OrderModel = (*customOrderModel)(nil)

type (
	// OrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderModel.
	OrderModel interface {
		orderModel
		FindAllByUid(ctx context.Context, uid int64) ([]*Order, error)
		TxInsert(tx *sql.Tx, data *Order) (sql.Result, error)
		TxUpdate(tx *sql.Tx, data *Order) error
		FindOneByUid(uid int64) (*Order, error)
		// 通过orderNo获取订单数据
		FindOneByOrderNo(orderNo string) (*Order, error)
	}

	customOrderModel struct {
		*defaultOrderModel
	}
)

// NewOrderModel returns a model for the database table.
func NewOrderModel(conn sqlx.SqlConn, c cache.CacheConf) OrderModel {
	return &customOrderModel{
		defaultOrderModel: newOrderModel(conn, c),
	}
}

func (m *customOrderModel) FindAllByUid(ctx context.Context, uid int64) ([]*Order, error) {
	var resp []*Order

	query := fmt.Sprintf("select %s from %s where `uid` = ?", orderRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, uid)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOrderModel) TxInsert(tx *sql.Tx, data *Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := tx.Exec(query, data.Uid, data.Pid, data.Amount, data.Status, data.OrderNo)

	return ret, err
}

func (m *customOrderModel) TxUpdate(tx *sql.Tx, data *Order) error {
	productIdKey := fmt.Sprintf("%s%v", cacheOrderIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
		return tx.Exec(query, data.Uid, data.Pid, data.Amount, data.Status, data.OrderNo, data.Id)
	}, productIdKey)
	return err
}

func (m *customOrderModel) FindOneByUid(uid int64) (*Order, error) {
	var resp Order

	query := fmt.Sprintf("select %s from %s where `uid` = ? order by create_time desc limit 1", orderRows, m.table)
	err := m.QueryRowNoCache(&resp, query, uid)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customOrderModel) FindOneByOrderNo(orderNo string) (*Order, error) {
	var resp Order

	query := fmt.Sprintf("select %s from %s where `order_no` = ? order by create_time desc limit 1", orderRows, m.table)
	err := m.QueryRowNoCache(&resp, query, orderNo)

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
