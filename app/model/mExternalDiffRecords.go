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
	_                                      MExternalDiffRecordsModel = (*defaultMExternalDiffRecordsModel)(nil)
	cacheMExternalDiffRecordsPrimaryPrefix                           = "cache:MExternalDiffRecords:primary:"
)

type (
	MExternalDiffRecords struct {
		Id              int64      `gorm:"id;primary_key"`    //
		ProjectId       int64      `gorm:"project_id"`        //
		TaskId          string     `gorm:"task_id"`           //
		StartArtifactId int64      `gorm:"start_artifact_id"` //
		StartCommit     string     `gorm:"start_commit"`      //
		StartTs         int64      `gorm:"start_ts"`          //
		EndArtifactId   int64      `gorm:"end_artifact_id"`   //
		EndCommit       string     `gorm:"end_commit"`        //
		EndTs           int64      `gorm:"end_ts"`            //
		Type            string     `gorm:"type"`              //
		LogFilePath     string     `gorm:"log_file_path"`     //
		DiffFilePath    string     `gorm:"diff_file_path"`    //
		Status          string     `gorm:"status"`            //
		CreatedAt       *time.Time `gorm:"created_at"`        //
		UpdatedAt       *time.Time `gorm:"updated_at"`        //
		DeletedAt       *time.Time `gorm:"deleted_at"`        //
	}

	MExternalDiffRecordsModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MExternalDiffRecords) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MExternalDiffRecords) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MExternalDiffRecordsQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MExternalDiffRecords, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MExternalDiffRecordsQuery) (list []*MExternalDiffRecords, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MExternalDiffRecordsQuery) (list []*MExternalDiffRecords, err error)
		FindAll(ctx context.Context, cond *MExternalDiffRecordsQuery) (list []*MExternalDiffRecords, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMExternalDiffRecordsModel struct {
		*customConn
		table string
	}
)

func NewMExternalDiffRecordsMode(db *gorm.DB, c cache.CacheConf) MExternalDiffRecordsModel {
	return &defaultMExternalDiffRecordsModel{
		customConn: newCustomConn(db, c),
		table:      "m_external_diff_records",
	}
}

func (m *defaultMExternalDiffRecordsModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMExternalDiffRecordsModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMExternalDiffRecordsModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMExternalDiffRecordsModel) Insert(ctx context.Context, db *gorm.DB, data *MExternalDiffRecords) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalDiffRecordsPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMExternalDiffRecordsModel) Update(ctx context.Context, db *gorm.DB, data *MExternalDiffRecords) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalDiffRecordsPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMExternalDiffRecordsModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalDiffRecordsPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMExternalDiffRecordsModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalDiffRecordsPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMExternalDiffRecordsModel) Count(ctx context.Context, cond *MExternalDiffRecordsQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MExternalDiffRecords).Count(&total).Error
	return total, err
}

func (m *defaultMExternalDiffRecordsModel) FindOne(ctx context.Context, Id int64) (data *MExternalDiffRecords, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMExternalDiffRecordsPrimaryPrefix, Id)
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

func (m *defaultMExternalDiffRecordsModel) FindListByPage(ctx context.Context, cond *MExternalDiffRecordsQuery) (list []*MExternalDiffRecords, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MExternalDiffRecords).Find(&list).Error
	return list, err
}

func (m *defaultMExternalDiffRecordsModel) FindListByCursor(ctx context.Context, cond *MExternalDiffRecordsQuery) (list []*MExternalDiffRecords, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MExternalDiffRecords).Find(&list).Error
	return list, err
}

func (m *defaultMExternalDiffRecordsModel) FindAll(ctx context.Context, cond *MExternalDiffRecordsQuery) (list []*MExternalDiffRecords, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MExternalDiffRecords).Find(&list).Error
	return list, err
}
