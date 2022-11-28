package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysTaskTemplateModel = (*customSysTaskTemplateModel)(nil)

type (
	// SysTaskTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysTaskTemplateModel.
	SysTaskTemplateModel interface {
		sysTaskTemplateModel
	}

	customSysTaskTemplateModel struct {
		*defaultSysTaskTemplateModel
	}
)

// NewSysTaskTemplateModel returns a model for the database table.
func NewSysTaskTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) SysTaskTemplateModel {
	return &customSysTaskTemplateModel{
		defaultSysTaskTemplateModel: newSysTaskTemplateModel(conn, c),
	}
}
