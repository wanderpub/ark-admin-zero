package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysTaskAccountModel = (*customSysTaskAccountModel)(nil)

type (
	// SysTaskAccountModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysTaskAccountModel.
	SysTaskAccountModel interface {
		sysTaskAccountModel
	}

	customSysTaskAccountModel struct {
		*defaultSysTaskAccountModel
	}
)

// NewSysTaskAccountModel returns a model for the database table.
func NewSysTaskAccountModel(conn sqlx.SqlConn, c cache.CacheConf) SysTaskAccountModel {
	return &customSysTaskAccountModel{
		defaultSysTaskAccountModel: newSysTaskAccountModel(conn, c),
	}
}
