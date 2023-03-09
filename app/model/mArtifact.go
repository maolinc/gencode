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
	_                           MArtifactModel = (*defaultMArtifactModel)(nil)
	cacheMArtifactPrimaryPrefix                = "cache:MArtifact:primary:"
)

type (
	MArtifact struct {
		Id             int64      `gorm:"id;primary_key"`   //制品id
		Name           string     `gorm:"name"`             //制品名称
		Source         int64      `gorm:"source"`           //制品来源(1-客户端，2-服务端)
		Version        string     `gorm:"version"`          //制品版本信息
		CommitId       string     `gorm:"commit_id"`        //git提交id
		CommitBranch   string     `gorm:"commit_branch"`    //git提交分支
		CompleteTime   *time.Time `gorm:"complete_time"`    //制品制造完成时间
		DownloadAddr   string     `gorm:"download_addr"`    //下载地址
		QaStatus       int64      `gorm:"qa_status"`        //qa状态(1-未qa，2-qa中，3-qa失败，4-qa通过)
		PushStatus     int64      `gorm:"push_status"`      //发布状态(1-未发布，2-发布中，3-发布失败，4-发布成功)
		DetailInfo     string     `gorm:"detail_info"`      //jenkins输出详情
		CreateUser     string     `gorm:"create_user"`      //制品创建者
		DeleteState    int64      `gorm:"delete_state"`     //
		DeleteTime     *time.Time `gorm:"delete_time"`      //
		CreateTime     *time.Time `gorm:"create_time"`      //创建时间
		UpdateTime     *time.Time `gorm:"update_time"`      //更新时间
		Ext1           string     `gorm:"ext1"`             //扩展字段
		Ext2           string     `gorm:"ext2"`             //扩展字段
		Ext3           string     `gorm:"ext3"`             //扩展字段
		Ext4           string     `gorm:"ext4"`             //扩展字段
		ExtInfo        string     `gorm:"ext_info"`         //扩展信息
		ProjectId      int64      `gorm:"project_id"`       //所属项目id
		ArtifactTypeId int64      `gorm:"artifact_type_id"` //制品类型id
		EnvId          int64      `gorm:"env_id"`           //环境id
		EnvRecord      string     `gorm:"env_record"`       //环境变更记录
		Code           string     `gorm:"code"`             //不能重复
	}

	MArtifactModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MArtifact) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MArtifact) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MArtifactQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MArtifact, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MArtifactQuery) (list []*MArtifact, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MArtifactQuery) (list []*MArtifact, err error)
		FindAll(ctx context.Context, cond *MArtifactQuery) (list []*MArtifact, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMArtifactModel struct {
		*customConn
		table string
	}
)

func NewMArtifactMode(db *gorm.DB, c cache.CacheConf) MArtifactModel {
	return &defaultMArtifactModel{
		customConn: newCustomConn(db, c),
		table:      "m_artifact",
	}
}

func (m *defaultMArtifactModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMArtifactModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMArtifactModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMArtifactModel) Insert(ctx context.Context, db *gorm.DB, data *MArtifact) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMArtifactModel) Update(ctx context.Context, db *gorm.DB, data *MArtifact) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMArtifactModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMArtifactModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMArtifactModel) Count(ctx context.Context, cond *MArtifactQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MArtifact).Count(&total).Error
	return total, err
}

func (m *defaultMArtifactModel) FindOne(ctx context.Context, Id int64) (data *MArtifact, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMArtifactPrimaryPrefix, Id)
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

func (m *defaultMArtifactModel) FindListByPage(ctx context.Context, cond *MArtifactQuery) (list []*MArtifact, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MArtifact).Find(&list).Error
	return list, err
}

func (m *defaultMArtifactModel) FindListByCursor(ctx context.Context, cond *MArtifactQuery) (list []*MArtifact, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MArtifact).Find(&list).Error
	return list, err
}

func (m *defaultMArtifactModel) FindAll(ctx context.Context, cond *MArtifactQuery) (list []*MArtifact, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MArtifact).Find(&list).Error
	return list, err
}
