{{ $PrimaryField := toCamelWithStartLower .Primary.CamelName }}
package model

import (
	"context"
	"database/sql"
	{{if .IsCache -}}
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	{{end -}}
	"gorm.io/gorm"
	"time"
)

var (
	_   {{.CamelName}}Model = (*default{{.CamelName}}Model)(nil)
	{{if .IsCache}}cache{{.CamelName}}PrimaryPrefix = "cache:{{.CamelName}}:primary:"{{end}}
)

type (
    {{.CamelName}} struct {
    {{range  .Fields}}    {{.CamelName}}  {{.DataType}}  `gorm:"{{.Name}} {{- if .IsPrimary}};primary_key{{end}}"` //{{.Comment}}
    {{end}}}

    {{.CamelName}}Model interface {
    	// Trans Transaction
        Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
    	// Builder Custom assembly conditions
    	Builder(ctx context.Context) (b *gorm.DB)
    	Insert(ctx context.Context, db *gorm.DB, data *{{.CamelName}}) (err error)
    	Update(ctx context.Context, db *gorm.DB, data *{{.CamelName}}) (err error)
    	// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
    	Delete(ctx context.Context, db *gorm.DB, {{$PrimaryField}} {{.Primary.DataType}}) (err error)
    	// ForceDelete Physical deletion
    	ForceDelete(ctx context.Context, db *gorm.DB, {{$PrimaryField}} {{.Primary.DataType}}) (err error)
    	Count(ctx context.Context, cond *{{.CamelName}}Query) (total int64, err error)
    	FindOne(ctx context.Context, {{$PrimaryField}} {{.Primary.DataType}}) (data *{{.CamelName}}, err error)
    	// FindListByPage Normal pagination
    	FindListByPage(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error)
    	// FindListByCursor Cursor is required based on cursor paging
    	FindListByCursor(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error)
    	FindAll(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error)
    	// ---------------Write your other interfaces below---------------
    }

    default{{.CamelName}}Model struct {
    	*customConn
    	table string
    }
)

func New{{.CamelName}}Model(db *gorm.DB {{- if .IsCache}}, c cache.CacheConf{{end}}) {{.CamelName}}Model {
	return &default{{.CamelName}}Model{
		customConn: {{if .IsCache}}newCustomConn(db, c){{else}}newCustomConnNoCache(db){{end}},
		table:      "{{.Name}}",
	}
}

func (m *default{{.CamelName}}Model) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *default{{.CamelName}}Model) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *default{{.CamelName}}Model) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *default{{.CamelName}}Model) Insert(ctx context.Context, db *gorm.DB, data *{{.CamelName}}) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, data.{{$PrimaryField}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Update(ctx context.Context, db *gorm.DB, data *{{.CamelName}}) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, data.{{$PrimaryField}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Delete(ctx context.Context, db *gorm.DB, {{$PrimaryField}} {{.Primary.DataType}}) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, {{$PrimaryField}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete({{$PrimaryField}}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) ForceDelete(ctx context.Context, db *gorm.DB, {{$PrimaryField}} {{.Primary.DataType}}) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, {{$PrimaryField}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete({{$PrimaryField}}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Count(ctx context.Context, cond *{{.CamelName}}Query) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.{{.CamelName}}).Count(&total).Error
	return total, err
}

func (m *default{{.CamelName}}Model) FindOne(ctx context.Context, {{$PrimaryField}} {{.Primary.DataType}}) (data *{{.CamelName}}, err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, {{$PrimaryField}})
    {{end -}}
	err = m.QueryRow(ctx, &data, func(v interface{}) error {
		tx := m.conn(ctx, nil).Find(v, {{$PrimaryField}})
		if tx.RowsAffected == 0 {
			return sql.ErrNoRows
		}
		return tx.Error
	} {{- if .IsCache}}, cacheKey{{end}})
	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *default{{.CamelName}}Model) FindListByPage(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error
	return list, err
}

func (m *default{{.CamelName}}Model) FindListByCursor(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "{{.Primary.Name}}"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error
	return list, err
}

func (m *default{{.CamelName}}Model) FindAll(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error
	return list, err
}
