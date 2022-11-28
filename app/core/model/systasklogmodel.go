package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysTaskLogModel = (*customSysTaskLogModel)(nil)

type (
	// SysTaskLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysTaskLogModel.
	SysTaskLogModel interface {
		sysTaskLogModel
	}

	customSysTaskLogModel struct {
		*defaultSysTaskLogModel
	}
)

// NewSysTaskLogModel returns a model for the database table.
func NewSysTaskLogModel(conn sqlx.SqlConn, c cache.CacheConf) SysTaskLogModel {
	return &customSysTaskLogModel{
		defaultSysTaskLogModel: newSysTaskLogModel(conn, c),
	}
}
