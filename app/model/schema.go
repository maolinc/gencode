package model

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

var (
	singleFlights = syncx.NewSingleFlight()
	stats         = cache.NewStat("sqlc")
	ErrNotCache   = errors.New("cache connection not set")
)

type customConn struct {
	db    *gorm.DB
	cache cache.Cache
}

func newCustomConn(db *gorm.DB, c cache.CacheConf, opts ...cache.Option) *customConn {
	cc := cache.New(c, singleFlights, stats, sql.ErrNoRows, opts...)
	return &customConn{
		db:    db,
		cache: cc,
	}
}

func newCustomConnNoCache(db *gorm.DB) *customConn {
	return &customConn{
		db:    db,
		cache: nil,
	}
}

func (c *customConn) Exec(ctx context.Context, exec func() error, keys ...string) error {
	if err := exec(); err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}
	if c.cache == nil {
		return ErrNotCache
	}
	if err := c.cache.DelCtx(ctx, keys...); err != nil {
		return err
	}
	return nil
}

func (c *customConn) QueryRow(ctx context.Context, v interface{}, query func(v interface{}) error, keys ...string) error {
	if len(keys) == 0 {
		return query(v)
	}
	if c.cache == nil {
		return ErrNotCache
	}
	return c.cache.TakeCtx(ctx, v, keys[0], func(v interface{}) error {
		return query(v)
	})
}

// -----------gorm  common scope----------------------
// Normal pagination
func pageScope(pageCurrent, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pageCurrent - 1) * pageSize
		return db.Limit(pageSize).Offset(offset)
	}
}

// Paging based on cursor
func cursorScope(cursorValue any, cursorAsc bool, pageSize int, field string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if cursorAsc {
			db.Where(field+" > ?", cursorValue)
		} else {
			db.Where(field+" < ?", cursorValue)
		}
		return db.Order(clause.OrderByColumn{Column: clause.Column{Name: field}, Desc: !cursorAsc}).Limit(pageSize)
	}
}

// eg: "create_time asc", "id desc"
func orderScope(sorts ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(sorts) == 0 {
			return db
		}
		return db.Order(strings.Join(sorts, ","))
	}
}
