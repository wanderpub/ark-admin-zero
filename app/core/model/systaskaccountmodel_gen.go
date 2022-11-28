// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysTaskAccountFieldNames          = builder.RawFieldNames(&SysTaskAccount{})
	sysTaskAccountRows                = strings.Join(sysTaskAccountFieldNames, ",")
	sysTaskAccountRowsExpectAutoSet   = strings.Join(stringx.Remove(sysTaskAccountFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	sysTaskAccountRowsWithPlaceHolder = strings.Join(stringx.Remove(sysTaskAccountFieldNames, "`id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cacheArkAdminSysTaskAccountIdPrefix = "cache:arkAdmin:sysTaskAccount:id:"
)

type (
	sysTaskAccountModel interface {
		Insert(ctx context.Context, data *SysTaskAccount) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysTaskAccount, error)
		Update(ctx context.Context, data *SysTaskAccount) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysTaskAccountModel struct {
		sqlc.CachedConn
		table string
	}

	SysTaskAccount struct {
		Id         int64  `db:"id"`
		Title      string `db:"title"`       // 账号名称
		SendChanel string `db:"send_chanel"` // 发送渠道
		Config     string `db:"config"`      // 账户配置
	}
)

func newSysTaskAccountModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysTaskAccountModel {
	return &defaultSysTaskAccountModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_task_account`",
	}
}

func (m *defaultSysTaskAccountModel) Delete(ctx context.Context, id int64) error {
	arkAdminSysTaskAccountIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskAccountIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, arkAdminSysTaskAccountIdKey)
	return err
}

func (m *defaultSysTaskAccountModel) FindOne(ctx context.Context, id int64) (*SysTaskAccount, error) {
	arkAdminSysTaskAccountIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskAccountIdPrefix, id)
	var resp SysTaskAccount
	err := m.QueryRowCtx(ctx, &resp, arkAdminSysTaskAccountIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysTaskAccountRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysTaskAccountModel) Insert(ctx context.Context, data *SysTaskAccount) (sql.Result, error) {
	arkAdminSysTaskAccountIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskAccountIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, sysTaskAccountRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Title, data.SendChanel, data.Config)
	}, arkAdminSysTaskAccountIdKey)
	return ret, err
}

func (m *defaultSysTaskAccountModel) Update(ctx context.Context, data *SysTaskAccount) error {
	arkAdminSysTaskAccountIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskAccountIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysTaskAccountRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Title, data.SendChanel, data.Config, data.Id)
	}, arkAdminSysTaskAccountIdKey)
	return err
}

func (m *defaultSysTaskAccountModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminSysTaskAccountIdPrefix, primary)
}

func (m *defaultSysTaskAccountModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysTaskAccountRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysTaskAccountModel) tableName() string {
	return m.table
}
