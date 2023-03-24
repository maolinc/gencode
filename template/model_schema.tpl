package model

import (
	"bytes"
	"context"
	"database/sql"
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
		//k         = false
	)

	return func(db *gorm.DB) *gorm.DB {
		for _, group := range plus {
			//k = false
			subQ.Reset()
			for _, searchItem := range group.Group {
				if searchItem.Field == "" {
					continue
				}
				value, typ, operator := searchItem.Value, strings.ToLower(searchItem.Type), searchItem.Operator
				var (
					v     any
					err   error
					logic = getLogic(searchItem.Logic)
				)

				switch typ {
				case "number":
					v, err = strconv.ParseInt(value, 10, 64)
				case "date":
					if v1, err := strconv.ParseInt(value, 10, 64); err == nil {
						if t := time.Unix(v1, 0); err == nil {
							v = t.String()
						}
					}
					if t, err := time.Parse("2006-01-02 15:04:05", value); err == nil {
						v = t.String()
					}
				case "string":
					if operator == "包含" {
						v = fmt.Sprintf("%%%v%%", value)
					} else if operator == "不包含" {
						v = fmt.Sprintf("%%%v%%", value)
					}
					v = value
				case "stringarray":
					v = strings.Split(value, ",")
				case "numberarray":
					v, _ = toNumberArray(value)
				default:
					v = value
				}

				if err != nil || v == nil {
					continue
				}

				if subQ.Len() != 0 {
					fmt.Fprintf(&subQ, " %s ", logic)
				}
                if searchItem.Table == "" {
                    fmt.Fprintf(&subQ, "`%s`.`%s` %s ?", tableName, searchItem.Field, getOperator(operator, typ))
                } else {
                    fmt.Fprintf(&subQ, "`%s`.`%s` %s ?", searchItem.Table, searchItem.Field, getOperator(operator, typ))
                }
				queryArgs = append(queryArgs, v)
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
	logic = strings.ToLower(logic)
	if logic != "and" && logic != "or" {
		return "and"
	}
	return logic
}

func toNumberArray(strArr string) ([]int64, error) {
	arr := make([]int64, 0)
	for _, s := range strings.Split(strArr, ",") {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return arr, err
		}
		arr = append(arr, i)
	}
	return arr, nil
}

func getOperator(oper, typ string) string {
	switch oper {
	case "包含":
		if typ == "string" {
			return "like"
		}
		return "in"
	case "不包含":
		if typ == "string" {
			return "not like"
		}
		return "not in"

	}
	return oper
}