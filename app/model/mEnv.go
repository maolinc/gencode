package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"time"
)

var (
	_                      MEnvModel = (*defaultMEnvModel)(nil)
	cacheMEnvPrimaryPrefix           = "cache:MEnv:primary:"
)

type (
	MEnv struct {
		Id          int64      `gorm:"id;primary_key"` //
		CreateUser  int64      `gorm:"create_user"`    //创建者id
		CreateTime  *time.Time `gorm:"create_time"`    //
		UpdateTime  *time.Time `gorm:"update_time"`    //
		DeleteTime  *time.Time `gorm:"delete_time"`    //
		DeleteState int64      `gorm:"delete_state"`   //
		Name        string     `gorm:"name"`           //环境名称
		Status      int64      `gorm:"status"`         //环境状态 1启用 2停用
		Des         string     `gorm:"des"`            //环境描述
		Expire      int64      `gorm:"expire"`         //过期时间(时),-1永不过期
	}

	MEnvModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MEnv) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MEnv) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MEnvQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MEnv, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MEnvQuery) (list []*MEnv, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MEnvQuery) (list []*MEnv, err error)
		FindAll(ctx context.Context, cond *MEnvQuery) (list []*MEnv, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMEnvModel struct {
		*customConn
		table string
	}
)

func NewMEnvMode(db *gorm.DB, c cache.CacheConf) MEnvModel {
	return &defaultMEnvModel{
		customConn: newCustomConn(db, c),
		table:      "m_env",
	}
}

func (m *defaultMEnvModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMEnvModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMEnvModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMEnvModel) Insert(ctx context.Context, db *gorm.DB, data *MEnv) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMEnvPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMEnvModel) Update(ctx context.Context, db *gorm.DB, data *MEnv) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMEnvPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMEnvModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMEnvPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMEnvModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMEnvPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMEnvModel) Count(ctx context.Context, cond *MEnvQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MEnv).Count(&total).Error
	return total, err
}

func (m *defaultMEnvModel) FindOne(ctx context.Context, Id int64) (data *MEnv, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMEnvPrimaryPrefix, Id)
	err = m.QueryRow(ctx, &data, func(v interface{}) error {
		tx := m.conn(ctx, nil).Find(v, Id)
		if tx.RowsAffected == 0 {
			return sql.ErrNoRows
		}
		return tx.Error
	}, cacheKey)
	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultMEnvModel) FindListByPage(ctx context.Context, cond *MEnvQuery) (list []*MEnv, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MEnv).Find(&list).Error
	return list, err
}

func (m *defaultMEnvModel) FindListByCursor(ctx context.Context, cond *MEnvQuery) (list []*MEnv, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MEnv).Find(&list).Error
	return list, err
}

func (m *defaultMEnvModel) FindAll(ctx context.Context, cond *MEnvQuery) (list []*MEnv, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MEnv).Find(&list).Error
	return list, err
}
