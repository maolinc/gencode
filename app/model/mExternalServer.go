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
	_                                 MExternalServerModel = (*defaultMExternalServerModel)(nil)
	cacheMExternalServerPrimaryPrefix                      = "cache:MExternalServer:primary:"
)

type (
	MExternalServer struct {
		Id        int64      `gorm:"id;primary_key"` //
		Type      string     `gorm:"type"`           //
		Name      string     `gorm:"name"`           //
		Proj      string     `gorm:"proj"`           //
		Addr      string     `gorm:"addr"`           //
		Desc      string     `gorm:"desc"`           //
		ApiMap    string     `gorm:"api_map"`        //
		Attach    string     `gorm:"attach"`         //
		CreatedAt *time.Time `gorm:"created_at"`     //
		UpdatedAt *time.Time `gorm:"updated_at"`     //
		DeletedAt *time.Time `gorm:"deleted_at"`     //
	}

	MExternalServerModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MExternalServer) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MExternalServer) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MExternalServerQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MExternalServer, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MExternalServerQuery) (list []*MExternalServer, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MExternalServerQuery) (list []*MExternalServer, err error)
		FindAll(ctx context.Context, cond *MExternalServerQuery) (list []*MExternalServer, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMExternalServerModel struct {
		*customConn
		table string
	}
)

func NewMExternalServerMode(db *gorm.DB, c cache.CacheConf) MExternalServerModel {
	return &defaultMExternalServerModel{
		customConn: newCustomConn(db, c),
		table:      "m_external_server",
	}
}

func (m *defaultMExternalServerModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMExternalServerModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMExternalServerModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMExternalServerModel) Insert(ctx context.Context, db *gorm.DB, data *MExternalServer) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalServerPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMExternalServerModel) Update(ctx context.Context, db *gorm.DB, data *MExternalServer) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalServerPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMExternalServerModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalServerPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMExternalServerModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalServerPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMExternalServerModel) Count(ctx context.Context, cond *MExternalServerQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MExternalServer).Count(&total).Error
	return total, err
}

func (m *defaultMExternalServerModel) FindOne(ctx context.Context, Id int64) (data *MExternalServer, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalServerPrimaryPrefix, Id)
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

func (m *defaultMExternalServerModel) FindListByPage(ctx context.Context, cond *MExternalServerQuery) (list []*MExternalServer, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MExternalServer).Find(&list).Error
	return list, err
}

func (m *defaultMExternalServerModel) FindListByCursor(ctx context.Context, cond *MExternalServerQuery) (list []*MExternalServer, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MExternalServer).Find(&list).Error
	return list, err
}

func (m *defaultMExternalServerModel) FindAll(ctx context.Context, cond *MExternalServerQuery) (list []*MExternalServer, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MExternalServer).Find(&list).Error
	return list, err
}
