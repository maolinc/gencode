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
	_                      MBugModel = (*defaultMBugModel)(nil)
	cacheMBugPrimaryPrefix           = "cache:MBug:primary:"
)

type (
	MBug struct {
		Id              int64      `gorm:"id;primary_key"`    //
		ProjectId       int64      `gorm:"project_id"`        //所属项目id
		CreateUser      int64      `gorm:"create_user"`       //
		CreateTime      *time.Time `gorm:"create_time"`       //
		UpdateTime      *time.Time `gorm:"update_time"`       //
		DeleteTime      *time.Time `gorm:"delete_time"`       //
		DeleteState     int64      `gorm:"delete_state"`      //
		MasterVersionId int64      `gorm:"master_version_id"` //迭代版本号
		SourceKind      int64      `gorm:"source_kind"`       //bug来源（1制品、2版本、3其他）
		SourceId        int64      `gorm:"source_id"`         //来源id
		SourceName      string     `gorm:"source_name"`       //来源名字
		EnvId           int64      `gorm:"env_id"`            //环境id
		Feedback        string     `gorm:"feedback"`          //反馈者
		Kind            string     `gorm:"kind"`              //bug类型
		Os              string     `gorm:"os"`                //操作系统（all、pc、ios、android）
		Serious         string     `gorm:"serious"`           //严重程度1.2.3.4,越大越重
		Priority        string     `gorm:"priority"`          //优先级1.2.3.4,越大越高
		Title           string     `gorm:"title"`             //bug名称
		Des             string     `gorm:"des"`               //bug描述
		Status          string     `gorm:"status"`            //状态
		Solve           string     `gorm:"solve"`             //解决方案
		CommitTime      *time.Time `gorm:"commit_time"`       //bug提交提交时间
		Appear          string     `gorm:"appear"`            //发现版本
		TitleUrl        string     `gorm:"title_url"`         //标题超链接
		IdUrl           string     `gorm:"id_url"`            //id超链接
		UploadTime      *time.Time `gorm:"upload_time"`       //上传时间
	}

	MBugModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MBug) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MBug) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MBugQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MBug, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MBugQuery) (list []*MBug, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MBugQuery) (list []*MBug, err error)
		FindAll(ctx context.Context, cond *MBugQuery) (list []*MBug, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMBugModel struct {
		*customConn
		table string
	}
)

func NewMBugMode(db *gorm.DB, c cache.CacheConf) MBugModel {
	return &defaultMBugModel{
		customConn: newCustomConn(db, c),
		table:      "m_bug",
	}
}

func (m *defaultMBugModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMBugModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMBugModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMBugModel) Insert(ctx context.Context, db *gorm.DB, data *MBug) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMBugPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMBugModel) Update(ctx context.Context, db *gorm.DB, data *MBug) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMBugPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMBugModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMBugPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMBugModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMBugPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMBugModel) Count(ctx context.Context, cond *MBugQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MBug).Count(&total).Error
	return total, err
}

func (m *defaultMBugModel) FindOne(ctx context.Context, Id int64) (data *MBug, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMBugPrimaryPrefix, Id)
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

func (m *defaultMBugModel) FindListByPage(ctx context.Context, cond *MBugQuery) (list []*MBug, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MBug).Find(&list).Error
	return list, err
}

func (m *defaultMBugModel) FindListByCursor(ctx context.Context, cond *MBugQuery) (list []*MBug, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MBug).Find(&list).Error
	return list, err
}

func (m *defaultMBugModel) FindAll(ctx context.Context, cond *MBugQuery) (list []*MBug, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MBug).Find(&list).Error
	return list, err
}
