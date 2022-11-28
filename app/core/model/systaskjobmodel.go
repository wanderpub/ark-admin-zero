package model

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysTaskJobModel = (*customSysTaskJobModel)(nil)

type (
	// SysTaskJobModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysTaskJobModel.
	SysTaskJobModel interface {
		sysTaskJobModel
		//tableName为空默认当前表名，fields为空默认所有字段
		SelectBuilder(tableName, fields string) squirrel.SelectBuilder
		UpdateBuilder() squirrel.UpdateBuilder
		DeleteBuilder() squirrel.DeleteBuilder
		InsertBuilder() squirrel.InsertBuilder

		//执行SQL语句，返回执行结果
		Exec(ctx context.Context, sql string, values []interface{}) (sql.Result, error)
		Select(ctx context.Context, sql string, values []interface{}) ([]*SysTaskJob, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		GetOne(ctx context.Context, builder squirrel.SelectBuilder) (int64, error)
		GetRow(ctx context.Context, builder squirrel.SelectBuilder) (*SysTaskJob, error)
		GetAll(ctx context.Context, builder squirrel.SelectBuilder) ([]*SysTaskJob, error)
		FindAll(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64) ([]*SysTaskJob, error)
	}

	customSysTaskJobModel struct {
		*defaultSysTaskJobModel
	}
)

// NewSysTaskJobModel returns a model for the database table.
func NewSysTaskJobModel(conn sqlx.SqlConn, c cache.CacheConf) SysTaskJobModel {
	return &customSysTaskJobModel{
		defaultSysTaskJobModel: newSysTaskJobModel(conn, c),
	}
}

func (m *defaultSysTaskJobModel) SelectBuilder(tableName, fields string) squirrel.SelectBuilder {
	if tableName == "" {
		tableName = m.table
	}
	if fields == "" {
		return squirrel.Select(sysTaskJobRows).From(tableName)
	} else {
		return squirrel.Select(fields).From(tableName)
	}
}

func (m *defaultSysTaskJobModel) UpdateBuilder() squirrel.UpdateBuilder {
	return squirrel.Update(m.table)
}

func (m *defaultSysTaskJobModel) DeleteBuilder() squirrel.DeleteBuilder {
	return squirrel.Delete(m.table)
}

func (m *defaultSysTaskJobModel) InsertBuilder() squirrel.InsertBuilder {
	return squirrel.Insert(m.table)
}

func (m *defaultSysTaskJobModel) Exec(ctx context.Context, sql string, values []interface{}) (sql.Result, error) {
	return m.ExecNoCacheCtx(ctx, sql, values...)
}

// 事务
func (m *defaultSysTaskJobModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultSysTaskJobModel) GetOne(ctx context.Context, builder squirrel.SelectBuilder) (int64, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return 0, err
	}
	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultSysTaskJobModel) GetRow(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*SysTaskJob, error) {
	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp *SysTaskJob
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysTaskJobModel) GetAll(ctx context.Context, builder squirrel.SelectBuilder) ([]*SysTaskJob, error) {
	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*SysTaskJob
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysTaskJobModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64) ([]*SysTaskJob, error) {
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize
	query, values, err := builder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}
	var resp []*SysTaskJob
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultSysTaskJobModel) Select(ctx context.Context, sql string, values []interface{}) ([]*SysTaskJob, error) {
	var resp []*SysTaskJob
	err := m.QueryRowsNoCacheCtx(ctx, &resp, sql, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
