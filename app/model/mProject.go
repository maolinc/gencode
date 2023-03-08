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
	_                          MProjectModel = (*defaultMProjectModel)(nil)
	cacheMProjectPrimaryPrefix               = "cache:MProject:primary:"
)

type (
	MProject struct {
		Id          int64      `gorm:"id;primary_key"` //
		CreateUser  int64      `gorm:"create_user"`    //创建者id
		CreateTime  *time.Time `gorm:"create_time"`    //
		UpdateTime  *time.Time `gorm:"update_time"`    //
		DeleteTime  *time.Time `gorm:"delete_time"`    //
		DeleteState int64      `gorm:"delete_state"`   //
		Name        string     `gorm:"name"`           //项目名称
		Des         string     `gorm:"des"`            //
		Status      int64      `gorm:"status"`         //状态(1启用，2停用)
	}

	MProjectModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MProject) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MProject) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MProjectQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MProject, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MProjectQuery) (list []*MProject, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MProjectQuery) (list []*MProject, err error)
		FindAll(ctx context.Context, cond *MProjectQuery) (list []*MProject, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMProjectModel struct {
		*customConn
		table string
	}
)

func NewMProjectMode(db *gorm.DB, c cache.CacheConf) MProjectModel {
	return &defaultMProjectModel{
		customConn: newCustomConn(db, c),
		table:      "m_project",
	}
}

func (m *defaultMProjectModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMProjectModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMProjectModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMProjectModel) Insert(ctx context.Context, db *gorm.DB, data *MProject) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMProjectPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMProjectModel) Update(ctx context.Context, db *gorm.DB, data *MProject) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMProjectPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMProjectModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMProjectPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMProjectModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMProjectPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMProjectModel) Count(ctx context.Context, cond *MProjectQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MProject).Count(&total).Error
	return total, err
}

func (m *defaultMProjectModel) FindOne(ctx context.Context, Id int64) (data *MProject, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMProjectPrimaryPrefix, Id)
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

func (m *defaultMProjectModel) FindListByPage(ctx context.Context, cond *MProjectQuery) (list []*MProject, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MProject).Find(&list).Error
	return list, err
}

func (m *defaultMProjectModel) FindListByCursor(ctx context.Context, cond *MProjectQuery) (list []*MProject, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MProject).Find(&list).Error
	return list, err
}

func (m *defaultMProjectModel) FindAll(ctx context.Context, cond *MProjectQuery) (list []*MProject, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MProject).Find(&list).Error
	return list, err
}
