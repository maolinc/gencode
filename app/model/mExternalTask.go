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
	_                               MExternalTaskModel = (*defaultMExternalTaskModel)(nil)
	cacheMExternalTaskPrimaryPrefix                    = "cache:MExternalTask:primary:"
)

type (
	MExternalTask struct {
		Id        int64      `gorm:"id;primary_key"` //
		TaskId    string     `gorm:"taskId"`         //
		Type      string     `gorm:"type"`           //
		Code      int64      `gorm:"code"`           //
		Res       string     `gorm:"res"`            //
		Attach    string     `gorm:"attach"`         //
		CreatedAt *time.Time `gorm:"created_at"`     //
		UpdatedAt *time.Time `gorm:"updated_at"`     //
		DeletedAt *time.Time `gorm:"deleted_at"`     //
	}

	MExternalTaskModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MExternalTask) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MExternalTask) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MExternalTaskQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MExternalTask, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MExternalTaskQuery) (list []*MExternalTask, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MExternalTaskQuery) (list []*MExternalTask, err error)
		FindAll(ctx context.Context, cond *MExternalTaskQuery) (list []*MExternalTask, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMExternalTaskModel struct {
		*customConn
		table string
	}
)

func NewMExternalTaskMode(db *gorm.DB, c cache.CacheConf) MExternalTaskModel {
	return &defaultMExternalTaskModel{
		customConn: newCustomConn(db, c),
		table:      "m_external_task",
	}
}

func (m *defaultMExternalTaskModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMExternalTaskModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMExternalTaskModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMExternalTaskModel) Insert(ctx context.Context, db *gorm.DB, data *MExternalTask) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalTaskPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMExternalTaskModel) Update(ctx context.Context, db *gorm.DB, data *MExternalTask) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalTaskPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMExternalTaskModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalTaskPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMExternalTaskModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalTaskPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMExternalTaskModel) Count(ctx context.Context, cond *MExternalTaskQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MExternalTask).Count(&total).Error
	return total, err
}

func (m *defaultMExternalTaskModel) FindOne(ctx context.Context, Id int64) (data *MExternalTask, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalTaskPrimaryPrefix, Id)
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

func (m *defaultMExternalTaskModel) FindListByPage(ctx context.Context, cond *MExternalTaskQuery) (list []*MExternalTask, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MExternalTask).Find(&list).Error
	return list, err
}

func (m *defaultMExternalTaskModel) FindListByCursor(ctx context.Context, cond *MExternalTaskQuery) (list []*MExternalTask, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MExternalTask).Find(&list).Error
	return list, err
}

func (m *defaultMExternalTaskModel) FindAll(ctx context.Context, cond *MExternalTaskQuery) (list []*MExternalTask, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MExternalTask).Find(&list).Error
	return list, err
}
