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
	_                          MVersionModel = (*defaultMVersionModel)(nil)
	cacheMVersionPrimaryPrefix               = "cache:MVersion:primary:"
)

type (
	MVersion struct {
		Id          int64      `gorm:"id;primary_key"` //id
		VersionNum  string     `gorm:"version_num"`    //版本号
		ServerId    string     `gorm:"server_id"`      //服务器
		ImageId     string     `gorm:"image_id"`       //镜像id
		ApkId       int64      `gorm:"apk_id"`         //apk包id
		BundleId    string     `gorm:"bundle_id"`      //bundle_id
		DataId      string     `gorm:"data_id"`        //数据档id
		QaStatus    int64      `gorm:"qa_status"`      //qa状态(1-未qa，2-qa中，3-qa失败，4-qa通过)
		PushStatus  int64      `gorm:"push_status"`    //发布状态(1-未发布，2-发布中，3-发布失败，4-发布成功)
		BugInfo     string     `gorm:"bug_info"`       //bug信息
		Des         string     `gorm:"des"`            //描述
		ProjectId   int64      `gorm:"project_id"`     //所属项目id
		EnvId       int64      `gorm:"env_id"`         //环境id
		CreateUser  string     `gorm:"create_user"`    //制品创建者
		DeleteState int64      `gorm:"delete_state"`   //
		DeleteTime  *time.Time `gorm:"delete_time"`    //
		CreateTime  *time.Time `gorm:"create_time"`    //创建时间
		UpdateTime  *time.Time `gorm:"update_time"`    //更新时间
		Ext1        string     `gorm:"ext1"`           //扩展字段
		Ext2        string     `gorm:"ext2"`           //扩展字段
		Ext3        string     `gorm:"ext3"`           //扩展字段
		Ext4        string     `gorm:"ext4"`           //扩展字段
		Pid         int64      `gorm:"pid"`            //父版本id
		Name        string     `gorm:"name"`           //名称
	}

	MVersionModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MVersion) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MVersion) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MVersionQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MVersion, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MVersionQuery) (list []*MVersion, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MVersionQuery) (list []*MVersion, err error)
		FindAll(ctx context.Context, cond *MVersionQuery) (list []*MVersion, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMVersionModel struct {
		*customConn
		table string
	}
)

func NewMVersionMode(db *gorm.DB, c cache.CacheConf) MVersionModel {
	return &defaultMVersionModel{
		customConn: newCustomConn(db, c),
		table:      "m_version",
	}
}

func (m *defaultMVersionModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMVersionModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMVersionModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMVersionModel) Insert(ctx context.Context, db *gorm.DB, data *MVersion) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMVersionPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMVersionModel) Update(ctx context.Context, db *gorm.DB, data *MVersion) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMVersionPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMVersionModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMVersionPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMVersionModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMVersionPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMVersionModel) Count(ctx context.Context, cond *MVersionQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MVersion).Count(&total).Error
	return total, err
}

func (m *defaultMVersionModel) FindOne(ctx context.Context, Id int64) (data *MVersion, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMVersionPrimaryPrefix, Id)
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

func (m *defaultMVersionModel) FindListByPage(ctx context.Context, cond *MVersionQuery) (list []*MVersion, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MVersion).Find(&list).Error
	return list, err
}

func (m *defaultMVersionModel) FindListByCursor(ctx context.Context, cond *MVersionQuery) (list []*MVersion, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MVersion).Find(&list).Error
	return list, err
}

func (m *defaultMVersionModel) FindAll(ctx context.Context, cond *MVersionQuery) (list []*MVersion, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MVersion).Find(&list).Error
	return list, err
}
