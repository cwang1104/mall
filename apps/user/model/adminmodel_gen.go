// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	adminFieldNames          = builder.RawFieldNames(&Admin{})
	adminRows                = strings.Join(adminFieldNames, ",")
	adminRowsExpectAutoSet   = strings.Join(stringx.Remove(adminFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), ",")
	adminRowsWithPlaceHolder = strings.Join(stringx.Remove(adminFieldNames, "`id`", "`update_time`", "`create_at`", "`created_at`", "`create_time`", "`update_at`", "`updated_at`"), "=?,") + "=?"

	cacheAdminIdPrefix       = "cache:admin:id:"
	cacheAdminUserNamePrefix = "cache:admin:userName:"
)

type (
	adminModel interface {
		Insert(ctx context.Context, data *Admin) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Admin, error)
		FindOneByUserName(ctx context.Context, userName string) (*Admin, error)
		Update(ctx context.Context, data *Admin) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAdminModel struct {
		sqlc.CachedConn
		table string
	}

	Admin struct {
		Id         int64     `db:"id"`        // ID
		UserName   string    `db:"user_name"` // 用户名
		Password   string    `db:"password"`  // 密码
		Desc       string    `db:"desc"`
		Status     int64     `db:"status"`
		CreateTime time.Time `db:"create_time"` // 创建时间
	}
)

func newAdminModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultAdminModel {
	return &defaultAdminModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`admin`",
	}
}

func (m *defaultAdminModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, id)
	adminUserNameKey := fmt.Sprintf("%s%v", cacheAdminUserNamePrefix, data.UserName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, adminIdKey, adminUserNameKey)
	return err
}

func (m *defaultAdminModel) FindOne(ctx context.Context, id int64) (*Admin, error) {
	adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, id)
	var resp Admin
	err := m.QueryRowCtx(ctx, &resp, adminIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminRows, m.table)
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

func (m *defaultAdminModel) FindOneByUserName(ctx context.Context, userName string) (*Admin, error) {
	adminUserNameKey := fmt.Sprintf("%s%v", cacheAdminUserNamePrefix, userName)
	var resp Admin
	err := m.QueryRowIndexCtx(ctx, &resp, adminUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", adminRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userName); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminModel) Insert(ctx context.Context, data *Admin) (sql.Result, error) {
	adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, data.Id)
	adminUserNameKey := fmt.Sprintf("%s%v", cacheAdminUserNamePrefix, data.UserName)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, adminRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserName, data.Password, data.Desc, data.Status)
	}, adminIdKey, adminUserNameKey)
	return ret, err
}

func (m *defaultAdminModel) Update(ctx context.Context, newData *Admin) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	adminIdKey := fmt.Sprintf("%s%v", cacheAdminIdPrefix, data.Id)
	adminUserNameKey := fmt.Sprintf("%s%v", cacheAdminUserNamePrefix, data.UserName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserName, newData.Password, newData.Desc, newData.Status, newData.Id)
	}, adminIdKey, adminUserNameKey)
	return err
}

func (m *defaultAdminModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheAdminIdPrefix, primary)
}

func (m *defaultAdminModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAdminModel) tableName() string {
	return m.table
}
