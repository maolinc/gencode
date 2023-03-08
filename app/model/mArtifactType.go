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
	_                               MArtifactTypeModel = (*defaultMArtifactTypeModel)(nil)
	cacheMArtifactTypePrimaryPrefix                    = "cache:MArtifactType:primary:"
)

type (
	MArtifactType struct {
		Id          int64      `gorm:"id;primary_key"` //
		CreateUser  int64      `gorm:"create_user"`    //创建者id
		CreateTime  *time.Time `gorm:"create_time"`    //
		UpdateTime  *time.Time `gorm:"update_time"`    //
		DeleteTime  *time.Time `gorm:"delete_time"`    //
		DeleteState int64      `gorm:"delete_state"`   //
		Name        string     `gorm:"name"`           //分类名称
		Status      int64      `gorm:"status"`         //分类状态 1启用 2停用
		Des         string     `gorm:"des"`            //描述
	}

	MArtifactTypeModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MArtifactType) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MArtifactType) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MArtifactTypeQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MArtifactType, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MArtifactTypeQuery) (list []*MArtifactType, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MArtifactTypeQuery) (list []*MArtifactType, err error)
		FindAll(ctx context.Context, cond *MArtifactTypeQuery) (list []*MArtifactType, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMArtifactTypeModel struct {
		*customConn
		table string
	}
)

func NewMArtifactTypeMode(db *gorm.DB, c cache.CacheConf) MArtifactTypeModel {
	return &defaultMArtifactTypeModel{
		customConn: newCustomConn(db, c),
		table:      "m_artifact_type",
	}
}

func (m *defaultMArtifactTypeModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMArtifactTypeModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMArtifactTypeModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMArtifactTypeModel) Insert(ctx context.Context, db *gorm.DB, data *MArtifactType) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactTypePrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMArtifactTypeModel) Update(ctx context.Context, db *gorm.DB, data *MArtifactType) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactTypePrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMArtifactTypeModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactTypePrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMArtifactTypeModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactTypePrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMArtifactTypeModel) Count(ctx context.Context, cond *MArtifactTypeQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MArtifactType).Count(&total).Error
	return total, err
}

func (m *defaultMArtifactTypeModel) FindOne(ctx context.Context, Id int64) (data *MArtifactType, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactTypePrimaryPrefix, Id)
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

func (m *defaultMArtifactTypeModel) FindListByPage(ctx context.Context, cond *MArtifactTypeQuery) (list []*MArtifactType, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MArtifactType).Find(&list).Error
	return list, err
}

func (m *defaultMArtifactTypeModel) FindListByCursor(ctx context.Context, cond *MArtifactTypeQuery) (list []*MArtifactType, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MArtifactType).Find(&list).Error
	return list, err
}

func (m *defaultMArtifactTypeModel) FindAll(ctx context.Context, cond *MArtifactTypeQuery) (list []*MArtifactType, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MArtifactType).Find(&list).Error
	return list, err
}
