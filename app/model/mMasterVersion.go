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
	_                                MMasterVersionModel = (*defaultMMasterVersionModel)(nil)
	cacheMMasterVersionPrimaryPrefix                     = "cache:MMasterVersion:primary:"
)

type (
	MMasterVersion struct {
		Id          int64      `gorm:"id;primary_key"` //
		CreateUser  int64      `gorm:"create_user"`    //创建者id
		CreateTime  *time.Time `gorm:"create_time"`    //
		UpdateTime  *time.Time `gorm:"update_time"`    //
		DeleteTime  *time.Time `gorm:"delete_time"`    //
		DeleteState int64      `gorm:"delete_state"`   //
		ProjectId   int64      `gorm:"project_id"`     //关联项目id
		Name        string     `gorm:"name"`           //名称
		VersionNum  string     `gorm:"version_num"`    //版本代号
		Status      int64      `gorm:"status"`         //环境状态 1启用 2停用
		Des         string     `gorm:"des"`            //清单描述
	}

	MMasterVersionModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MMasterVersion) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MMasterVersion) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MMasterVersionQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MMasterVersion, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MMasterVersionQuery) (list []*MMasterVersion, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MMasterVersionQuery) (list []*MMasterVersion, err error)
		FindAll(ctx context.Context, cond *MMasterVersionQuery) (list []*MMasterVersion, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMMasterVersionModel struct {
		*customConn
		table string
	}
)

func NewMMasterVersionMode(db *gorm.DB, c cache.CacheConf) MMasterVersionModel {
	return &defaultMMasterVersionModel{
		customConn: newCustomConn(db, c),
		table:      "m_master_version",
	}
}

func (m *defaultMMasterVersionModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMMasterVersionModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMMasterVersionModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMMasterVersionModel) Insert(ctx context.Context, db *gorm.DB, data *MMasterVersion) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMMasterVersionPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMMasterVersionModel) Update(ctx context.Context, db *gorm.DB, data *MMasterVersion) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMMasterVersionPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMMasterVersionModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMMasterVersionPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMMasterVersionModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMMasterVersionPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMMasterVersionModel) Count(ctx context.Context, cond *MMasterVersionQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MMasterVersion).Count(&total).Error
	return total, err
}

func (m *defaultMMasterVersionModel) FindOne(ctx context.Context, Id int64) (data *MMasterVersion, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMMasterVersionPrimaryPrefix, Id)
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

func (m *defaultMMasterVersionModel) FindListByPage(ctx context.Context, cond *MMasterVersionQuery) (list []*MMasterVersion, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MMasterVersion).Find(&list).Error
	return list, err
}

func (m *defaultMMasterVersionModel) FindListByCursor(ctx context.Context, cond *MMasterVersionQuery) (list []*MMasterVersion, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MMasterVersion).Find(&list).Error
	return list, err
}

func (m *defaultMMasterVersionModel) FindAll(ctx context.Context, cond *MMasterVersionQuery) (list []*MMasterVersion, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MMasterVersion).Find(&list).Error
	return list, err
}
