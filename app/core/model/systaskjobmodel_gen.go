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
	sysTaskJobFieldNames          = builder.RawFieldNames(&SysTaskJob{})
	sysTaskJobRows                = strings.Join(sysTaskJobFieldNames, ",")
	sysTaskJobRowsExpectAutoSet   = strings.Join(stringx.Remove(sysTaskJobFieldNames, "`job_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), ",")
	sysTaskJobRowsWithPlaceHolder = strings.Join(stringx.Remove(sysTaskJobFieldNames, "`job_id`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`", "`update_time`"), "=?,") + "=?"

	cacheArkAdminSysTaskJobJobIdPrefix = "cache:arkAdmin:sysTaskJob:jobId:"
)

type (
	sysTaskJobModel interface {
		Insert(ctx context.Context, data *SysTaskJob) (sql.Result, error)
		FindOne(ctx context.Context, jobId int64) (*SysTaskJob, error)
		Update(ctx context.Context, data *SysTaskJob) error
		Delete(ctx context.Context, jobId int64) error
	}

	defaultSysTaskJobModel struct {
		sqlc.CachedConn
		table string
	}

	SysTaskJob struct {
		JobId          int64  `db:"job_id"`
		JobName        string `db:"job_name"`
		JobGroup       string `db:"job_group"`
		JobType        int64  `db:"job_type"`
		CronExpression string `db:"cron_expression"`
		InvokeTarget   string `db:"invoke_target"`
		Args           string `db:"args"`
		MisfirePolicy  int64  `db:"misfire_policy"`
		Concurrent     int64  `db:"concurrent"`
		Status         int64  `db:"status"`
		EntryId        string `db:"entry_id"`
		CreateBy       int64  `db:"create_by"` // 创建时间
		UpdateBy       int64  `db:"update_by"` // 更新时间
		CreatedAt      int64  `db:"created_at"`
		UpdatedAt      int64  `db:"updated_at"`
		DeletedAt      int64  `db:"deleted_at"`
	}
)

func newSysTaskJobModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysTaskJobModel {
	return &defaultSysTaskJobModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`sys_task_job`",
	}
}

func (m *defaultSysTaskJobModel) Delete(ctx context.Context, jobId int64) error {
	arkAdminSysTaskJobJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskJobJobIdPrefix, jobId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `job_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, jobId)
	}, arkAdminSysTaskJobJobIdKey)
	return err
}

func (m *defaultSysTaskJobModel) FindOne(ctx context.Context, jobId int64) (*SysTaskJob, error) {
	arkAdminSysTaskJobJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskJobJobIdPrefix, jobId)
	var resp SysTaskJob
	err := m.QueryRowCtx(ctx, &resp, arkAdminSysTaskJobJobIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `job_id` = ? limit 1", sysTaskJobRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, jobId)
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

func (m *defaultSysTaskJobModel) Insert(ctx context.Context, data *SysTaskJob) (sql.Result, error) {
	arkAdminSysTaskJobJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskJobJobIdPrefix, data.JobId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysTaskJobRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.JobName, data.JobGroup, data.JobType, data.CronExpression, data.InvokeTarget, data.Args, data.MisfirePolicy, data.Concurrent, data.Status, data.EntryId, data.CreateBy, data.UpdateBy, data.DeletedAt)
	}, arkAdminSysTaskJobJobIdKey)
	return ret, err
}

func (m *defaultSysTaskJobModel) Update(ctx context.Context, data *SysTaskJob) error {
	arkAdminSysTaskJobJobIdKey := fmt.Sprintf("%s%v", cacheArkAdminSysTaskJobJobIdPrefix, data.JobId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `job_id` = ?", m.table, sysTaskJobRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.JobName, data.JobGroup, data.JobType, data.CronExpression, data.InvokeTarget, data.Args, data.MisfirePolicy, data.Concurrent, data.Status, data.EntryId, data.CreateBy, data.UpdateBy, data.DeletedAt, data.JobId)
	}, arkAdminSysTaskJobJobIdKey)
	return err
}

func (m *defaultSysTaskJobModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheArkAdminSysTaskJobJobIdPrefix, primary)
}

func (m *defaultSysTaskJobModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `job_id` = ? limit 1", sysTaskJobRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysTaskJobModel) tableName() string {
	return m.table
}
