package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPermMenuModel = (*customSysPermMenuModel)(nil)

type (
	// SysPermMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPermMenuModel.
	SysPermMenuModel interface {
		sysPermMenuModel
		FindByIds(ctx context.Context, ids string) ([]*SysPermMenu, error)
		FindAll(ctx context.Context) ([]*SysPermMenu, error)
	}

	customSysPermMenuModel struct {
		*defaultSysPermMenuModel
	}
)

// NewSysPermMenuModel returns a model for the database table.
func NewSysPermMenuModel(conn sqlx.SqlConn, c cache.CacheConf) SysPermMenuModel {
	return &customSysPermMenuModel{
		defaultSysPermMenuModel: newSysPermMenuModel(conn, c),
	}
}

func (m *customSysPermMenuModel) FindByIds(ctx context.Context, ids string) ([]*SysPermMenu, error) {
	query := fmt.Sprintf("select %s from %s where `id` in(%s)", sysPermMenuRows, m.table, ids)
	var resp []*SysPermMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customSysPermMenuModel) FindAll(ctx context.Context) ([]*SysPermMenu, error) {
	query := fmt.Sprintf("select %s from %s", sysPermMenuRows, m.table)
	var resp []*SysPermMenu
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}