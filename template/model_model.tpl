{{ $PrimaryField := toCamelWithStartLower .Primary.CamelName }}
package model

import (
	"context"
	{{if .IsCache}}"fmt"{{end -}}
	{{if .IsCache}}"github.com/zeromicro/go-zero/core/stores/cache"{{end -}}
	"gorm.io/gorm"
)

var (
	_   custom{{.CamelName}}Model = (*{{.CamelName}}Model)(nil)
	{{if .IsCache}}cache{{.CamelName}}PrimaryPrefix = "cache:{{.CamelName}}:primary:"{{end}}
)

type (
    {{.CamelName}} struct {
    {{range  .Fields}}    {{.CamelName}}  {{.DataType}}  `gorm:"{{.Name}} {{- if .IsPrimary}};primary_key{{end}}"` //{{.Comment}}
    {{end}}}

    custom{{.CamelName}}Model interface {
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

    {{.CamelName}}Model struct {
    	*customConn
    	table string
    }
)

func New{{.CamelName}}Mode(db *gorm.DB {{- if .IsCache}}, c cache.CacheConf{{end}}) *{{.CamelName}}Model {
	return &{{.CamelName}}Model{
		customConn: {{if .IsCache}}newCustomConn(db, c){{else}}newCustomConnNoCache(db){{end}},
		table:      "{{.Name}}",
	}
}

func (m *{{.CamelName}}Model) conn(db *gorm.DB, ctx context.Context) *gorm.DB {
	if db != nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *{{.CamelName}}Model) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(nil, ctx)
}

func (m *{{.CamelName}}Model) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(nil, ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *{{.CamelName}}Model) Insert(ctx context.Context, db *gorm.DB, data *{{.CamelName}}) (err error) {
	{{if .IsCache}}cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, data.{{$PrimaryField}}){{end -}}

	return m.Exec(ctx, func() error {
		return m.conn(db, ctx).Create(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *{{.CamelName}}Model) Update(ctx context.Context, db *gorm.DB, data *{{.CamelName}}) (err error) {
	{{if .IsCache}}cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, data.{{$PrimaryField}}){{end -}}

	return m.Exec(ctx, func() error {
		return m.conn(db, ctx).Updates(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *{{.CamelName}}Model) Delete(ctx context.Context, db *gorm.DB, {{$PrimaryField}} {{.Primary.DataType}}) (err error) {
	{{if .IsCache}}cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, {{$PrimaryField}}){{end -}}

	return m.Exec(ctx, func() error {
		return m.conn(db, ctx).Delete({{$PrimaryField}}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *{{.CamelName}}Model) ForceDelete(ctx context.Context, db *gorm.DB, {{$PrimaryField}} {{.Primary.DataType}}) (err error) {
	{{if .IsCache}}cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, {{$PrimaryField}}){{end -}}

	return m.Exec(ctx, func() error {
		return m.conn(db, ctx).Unscoped().Delete({{$PrimaryField}}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *{{.CamelName}}Model) Count(ctx context.Context, cond *{{.CamelName}}Query) (total int64, err error) {
	err = m.conn(nil, ctx).Where(cond.{{.CamelName}}).Count(&total).Error
	return total, err
}

func (m *{{.CamelName}}Model) FindOne(ctx context.Context, {{$PrimaryField}} {{.Primary.DataType}}) (data *{{.CamelName}}, err error) {
	{{if .IsCache}}cacheKey := fmt.Sprintf("%s%v", cache{{.CamelName}}PrimaryPrefix, {{$PrimaryField}}){{end -}}
	var records int64

	err = m.Exec(ctx, func() error {
		tx := m.conn(nil, ctx).Find(&data, {{$PrimaryField}})
		records = tx.RowsAffected
		return tx.Error
	} {{- if .IsCache}}, cacheKey{{end}})

	if err != nil {
		return nil, err
	}
	if records == 0 {
		return nil, nil
	}

	return data, nil
}

func (m *{{.CamelName}}Model) FindListByPage(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(nil, ctx)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error

	return list, err
}

func (m *{{.CamelName}}Model) FindListByCursor(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(nil, ctx)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "{{.Primary.Name}}"),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error

	return list, err
}

func (m *{{.CamelName}}Model) FindAll(ctx context.Context, cond *{{.CamelName}}Query) (list []*{{.CamelName}}, err error) {
	conn := m.conn(nil, ctx)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.{{.CamelName}}).Find(&list).Error

	return list, err
}
