
package model

import (
	"context"
	"database/sql"
	{{if .IsCache -}}
	"fmt"
	"time"
    "gitee.com/maolinc/vision-soul-common/collectx"
	"gitee.com/maolinc/vision-soul-common/utilx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

    // {{.CamelName}}Query query cond
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
    	// FindListByIds
    	FindListByIds(ctx context.Context, ids []{{.Primary.DataType}}) (list []*{{.CamelName}}, err error)
    	// ---------------Write your other interfaces below---------------
    }

    default{{.CamelName}}Model struct {
    	*customConn
    	{{if .IsCache}}
		redis  *redis.Redis
		expire time.Duration
        {{end}}
    	table string
    }
)

func New{{.CamelName}}Model(db *gorm.DB {{- if .IsCache}}, c cache.CacheConf, redis *redis.Redis{{end}}) {{.CamelName}}Model {
	return &default{{.CamelName}}Model{
		customConn: {{if .IsCache}}newCustomConn(db, c){{else}}newCustomConnNoCache(db){{end}},
		{{if .IsCache}}
        redis:      redis,
        expire:     10,
		{{end}}
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
	cacheKey := m.getPrimaryCacheKey(data.{{.Primary.CamelName}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Create(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Update(ctx context.Context, data *{{.CamelName}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := m.getPrimaryCacheKey(data.{{.Primary.CamelName}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Model(&data).Updates(data).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) Delete(ctx context.Context, {{.PrimaryFields}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := m.getPrimaryCacheKey({{.PrimaryFmtV2}})
	{{end -}}
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db...).Where("{{.PrimaryFieldWhere}}", {{.PrimaryFmtV2}}).Delete({{.CamelName}}{}).Error
	} {{- if .IsCache}}, cacheKey{{end}})
}

func (m *default{{.CamelName}}Model) ForceDelete(ctx context.Context, {{.PrimaryFields}}, db ...*gorm.DB) (err error) {
	{{if .IsCache -}}
	cacheKey := m.getPrimaryCacheKey({{.PrimaryFmtV2}})
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
	cacheKey := m.getPrimaryCacheKey({{.PrimaryFmtV2}})
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
	).Where(cond.{{.CamelName}})

	{{if .IsCache}}
    ids := make([]{{.Primary.DataType}}, 0, cond.PageSize)
    if err = conn.Pluck("{{.Primary.Name}}", &ids).Error; err != nil {
        return nil, err
    }
    if len(ids) == 0 {
        return list, nil
    }

    list, err = m.FindListByIds(ctx, ids)
	{{else}}
	err = conn.Find(&list).Error
	{{end}}

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

func (m *default{{.CamelName}}Model) FindListByIds(ctx context.Context, ids []{{.Primary.DataType}}) (list []*{{.CamelName}}, err error) {
    {{if .IsCache}}
    idKeys := m.getCacheKeysByIds(ids)
    jsonArr, err := m.redis.Mget(idKeys...)
    if err != nil {
        return list, nil
    }
    list = utilx.Json2StructArr[{{.CamelName}}](jsonArr)
    if len(list) == len(ids) {
        return list, nil
    }

    notExistIds := make([]{{.Primary.DataType}}, 0)
    set := collectx.NewSet[int64]()
    for _, item := range list {
        set.Add(item.{{.Primary.CamelName}})
    }
    for _, id := range ids {
        if set.Add(id) {
            notExistIds = append(notExistIds, id)
        }
    }
    dbList := make([]*{{.CamelName}}, 0, len(notExistIds))
    err = m.conn(ctx).Where("{{.Primary.Name}} in ?", notExistIds).Find(&dbList).Error
    if err != nil {
        return nil, err
    }

    _ = m.redis.PipelinedCtx(ctx, func(pipeliner redis.Pipeliner) error {
        for _, item := range dbList {
            key := m.getPrimaryCacheKey(item.{{.Primary.CamelName}})
            pipeliner.SetEX(ctx, key, item, m.expire)
        }
        return nil
    })

    return append(list, dbList...), nil
    {{else}}
    {{end}}
}


{{if .IsCache}}
func (m *default{{.CamelName}}Model) getPrimaryCacheKey({{.PrimaryFields}}) string {
	return fmt.Sprintf("%s{{.PrimaryFmt}}", cache{{.CamelName}}PrimaryPrefix, {{.PrimaryFmtV2}})
}

func (m *default{{.CamelName}}Model) getCacheKeysByIds(ids []{{.Primary.DataType}}) []string {
	idKeys := make([]string, 0, len(ids))
	for _, id := range ids {
		idKeys = append(idKeys, fmt.Sprintf("%s%v", cachePicturePrimaryPrefix, id))
	}
	return idKeys
}

{{end}}