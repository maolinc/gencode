package model

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/syncx"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"strings"
	"time"
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

func getCacheKeysByIds(prefixFormat string, ids []any) []string {
	idKeys := make([]string, 0, len(ids))
	for _, id := range ids {
		idKeys = append(idKeys, fmt.Sprintf("%s%v", prefixFormat, id))
	}
	return idKeys
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

func pageHandler[T any](conn *gorm.DB, currentPage, sizePage int) (total int64, list []T, err error) {
	//pC := make(chan int64)
	//go func(con *gorm.DB) {
	//	var count int64
	//	con.Count(&count)
	//	pC <- count
	//}(conn)
	//nConn := conn.Session(&gorm.Session{})
	//err = nConn.Scopes(pageScope(currentPage, sizePage)).Find(&list).Error
	//total = <-pC
	conn.Count(&total)
	err = conn.Scopes(pageScope(currentPage, sizePage)).Find(&list).Error
	return total, list, err
}

/*
[

	{
		"group": [
			{"field": "name","value": "没物表量","type": "string","operator": "=","logic": "and"},
			{"field": "name","value": "3","type": "string","operator": "不包含","logic": "or"}
		],
		"logic": "",
	},
	{
		"group": [
			{"field": "qa_status","value": "[1,3,4]","type": "stringArray","operator": "包含","logic": "and"},
			{"field": "create_time","value": "5522113322","type": "date","operator": ">","logic": "or"}
		],
		"logic": "",
	}

]
*/
func searchPlusScope(plus []SearchGroup, tableName string) func(db *gorm.DB) *gorm.DB {
	var (
		query     = bytes.Buffer{}
		subQ      = bytes.Buffer{}
		queryArgs = make([]any, 0)
		h         = false
	)

	return func(db *gorm.DB) *gorm.DB {
		for _, group := range plus {
			subQ.Reset()
			for _, searchItem := range group.Group {
				if searchItem.Field == "" {
					continue
				}
				typ, operator := strings.ToLower(searchItem.Type), strings.ToLower(searchItem.Operator)
				var (
					err   error
					logic = getLogic(searchItem.Logic)
				)
				value, err := getRealValue(typ, searchItem.Value)
				if err != nil || value == nil {
					continue
				}

				switch vv := reflect.ValueOf(value); vv.Kind() {
				case reflect.Slice:
					if operator == "不包含" || operator == "!=" {
						operator = "NOT IN"
					} else if operator == "包含" || operator == "=" {
						operator = "IN"
					}
				default:
					if operator == "不包含" {
						operator = "NOT LIKE"
					} else if operator == "包含" {
						operator = "LIKE"
					}
					value = fmt.Sprintf("%%%v%%", value)
				}

				if subQ.Len() != 0 {
					fmt.Fprintf(&subQ, " %s ", logic)
				}
				if searchItem.Table == "" {
					fmt.Fprintf(&subQ, "`%s`.`%s` %s ?", tableName, searchItem.Field, operator)
				} else {
					fmt.Fprintf(&subQ, "`%s`.`%s` %s ?", searchItem.Table, searchItem.Field, operator)
				}
				queryArgs = append(queryArgs, value)
			}
			if subQ.Len() == 0 {
				continue
			}
			if query.Len() != 0 {
				h = true
				fmt.Fprintf(&query, " %s ", getLogic(group.Logic))
			}
			query.WriteString("( ")
			query.Write(subQ.Bytes())
			query.WriteString(" )")
		}
		if !h && query.Len() != 0 {
			b := query.Bytes()
			return db.Where(string(b[2:len(b)-2]), queryArgs...)
		}

		return db.Where(query.String(), queryArgs...)
	}
}

func getLogic(logic string) string {
	logic = strings.ToUpper(logic)
	if logic != "And" && logic != "OR" {
		return "And"
	}
	return logic
}

func getRealValue(typ string, value interface{}) (interface{}, error) {
	var (
		v   interface{}
		err error
	)
	switch typ {
	case "date":
		var dataT time.Time
		switch vv := value.(type) {
		case float64:
			dataT = time.Unix(int64(vv), 0)
		case float32:
			dataT = time.Unix(int64(vv), 0)
		case int64:
			dataT = time.Unix(vv, 0)
		case int:
			dataT = time.Unix(int64(vv), 0)
		case string:
			if dataT, err = time.Parse("2006-01-02 15:04:05", vv); err != nil {
				dataT = time.Now()
			}
		}
		v = dataT.UTC().Format("2006-01-02 15:04:05")
	default:
		v = value
	}

	return v, err
}
