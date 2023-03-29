
package model

import (
	"context"
	"database/sql"
	{{if .IsCache -}}
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	{{end -}}
	"gorm.io/gorm"
	{{if .IsDate}}"time"{{end -}}
)

var (
	_   {{.CamelName}}Model = (*default{{.CamelName}}Model)(nil)
	{{if .IsCache}}cache{{.CamelName}}PrimaryPrefix = "cache:{{.CamelName}}:primary:"{{end}}
)

type (
    {{.CamelName}} struct {
    {{range  .Fields}}    {{.CamelName}}  {{.DataType}}  `gorm:"{{.Name}} {{- if .IsPrimary}};primary_key{{end}}"` //{{.Comment}}
    {{end}}}

    // {{.CamelName}} query cond
    {{.CamelName}}Query struct {
        SearchBase
        {{.CamelName}}
    }

    {{.CamelName}}Model interface {
    	// Trans Transaction
        Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
    	// Builder Custom assembly conditions
    	Builder(ctx context.Context, db ...*gorm.DB) (b *gorm.DB)
    	Insert(ctx context.Context, data *{{.CamelName}}, db ...*gorm.DB) (err error)
    	Update(ctx context.Context, data *{{.CamelName}}, db ...*gorm.DB) (err error)
    	// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
    	Delete(ctx context.Context, {{.PrimaryFields}}, db ...*gorm.DB) (err error)
    	// ForceDelete Physical deletion
    	ForceDelete(ctx context.Context, {{.PrimaryFields}}, db ...*gorm.DB) (err error)
    	Count(ctx context.Context, cond *{{.CamelName}}Query) (total int64, err error)
    	FindOne(ctx context.Context, {{.PrimaryFields}}) (data *{{.CamelName}}, err error)
    	// FindByPage Contains total information
    	FindByPage(ctx context.Context, cond *{{.CamelName}}Query) (total int64, list []*{{.CamelName}}, err error)
    	// FindListByPage Normal pagination
    	FindListByPage(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error)
    	// FindListByCursor Cursor is required based on cursor paging, Only the primary key is of type int, and other types can be expanded by themselves
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

func (m *{{.CamelName}}) TableName() string {
	return "`{{.Name}}`"
}

func (m *default{{.CamelName}}Model) conn(ctx context.Context, db ...*gorm.DB) *gorm.DB {
	if len(db) == 0 {
		return m.db.Model(&{{.CamelName}}{}).Session(&gorm.Session{Context: ctx})
	}
	return db[0]
}

func (m *default{{.CamelName}}Model) Builder(ctx context.Context, db ...*gorm.DB) (b *gorm.DB) {
	return m.conn(ctx, db...)
}

func (m *default{{.CamelName}}Model) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *default{{.CamelName}}Model) Insert(ctx context.Context, data *{{.CamelName}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s{{.PrimaryFmt}}", cache{{.CamelName}}PrimaryPrefix, {{.PrimaryFmtV}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Create(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Update(ctx context.Context, data *{{.CamelName}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s{{.PrimaryFmt}}", cache{{.CamelName}}PrimaryPrefix, {{.PrimaryFmtV}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Model(&data).Updates(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Delete(ctx context.Context, {{.PrimaryFields}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s{{.PrimaryFmt}}", cache{{.CamelName}}PrimaryPrefix, {{.PrimaryFmtV2}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Where("{{.PrimaryFieldWhere}}", {{.PrimaryFmtV2}}).Delete({{.CamelName}}{}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) ForceDelete(ctx context.Context, {{.PrimaryFields}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s{{.PrimaryFmt}}", cache{{.CamelName}}PrimaryPrefix, {{.PrimaryFmtV2}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Unscoped().Where("{{.PrimaryFieldWhere}}", {{.PrimaryFmtV2}}).Delete({{.CamelName}}{}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Count(ctx context.Context, cond *{{.CamelName}}Query) (total int64, err error) {
    err = m.conn(ctx).Scopes(
        searchPlusScope(cond.SearchPlus, m.table),
    ).Where(cond.{{.CamelName}}).Count(&total).Error
    return total, err
}

func (m *default{{.CamelName}}Model) FindOne(ctx context.Context, {{.PrimaryFields}}) (data *{{.CamelName}}, err error) {
	{{if .IsCache -}}
	cacheKey := fmt.Sprintf("%s{{.PrimaryFmt}}", cache{{.CamelName}}PrimaryPrefix, {{.PrimaryFmtV2}})
    {{end -}}
	err = m.QueryRow(ctx, &data, func(v interface{}) error {
		tx := m.conn(ctx).Where("{{.PrimaryFieldWhere}}", {{.PrimaryFmtV2}}).Find(v)
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

func (m *default{{.CamelName}}Model) FindByPage(ctx context.Context, cond *{{.CamelName}}Query) (total int64, list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx).Scopes(
		searchPlusScope(cond.SearchPlus, m.table),
	).Where(cond.{{.CamelName}})

	total, list, err = pageHandler[*{{.CamelName}}](conn, cond.PageCurrent, cond.PageSize)
	return total, list, err
}

func (m *default{{.CamelName}}Model) FindListByPage(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx).Scopes(
	    searchPlusScope(cond.SearchPlus, m.table),
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error
	return list, err
}

func (m *default{{.CamelName}}Model) FindListByCursor(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx).Scopes(
	    searchPlusScope(cond.SearchPlus, m.table),
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "{{.Primary.Name}}"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error
	return list, err
}

func (m *default{{.CamelName}}Model) FindAll(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(ctx).Scopes(
	    searchPlusScope(cond.SearchPlus, m.table),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error
	return list, err
}
