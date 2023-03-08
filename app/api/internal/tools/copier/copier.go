package copier

import (
	"errors"
	"github.com/jinzhu/copier"
	"strconv"
	"time"
)

var (
	NotMatchErr        = errors.New("src type not matching")
	FmtDateTime        = "2006-01-02 15:04:05"
	Int64       int64  = 0
	String      string = ""
	// 类型转换规则
	tcs = []copier.TypeConverter{
		{
			SrcType: time.Time{},
			DstType: Int64,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(time.Time); ok {
					return s.Unix(), nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: time.Time{},
			DstType: String,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(time.Time); ok {
					return s.Format(FmtDateTime), nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: Int64,
			DstType: time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(int64); ok {
					time := time.Unix(s, 0)
					return time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: String,
			DstType: time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(string); ok {
					time, _ := time.Parse(FmtDateTime, s)
					return time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: Int64,
			DstType: &time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(int64); ok {
					time := time.Unix(s, 0)
					return &time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: String,
			DstType: &time.Time{},
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(string); ok {
					time, _ := time.Parse(FmtDateTime, s)
					return &time, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: String,
			DstType: Int64,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(string); ok {
					i, _ := strconv.ParseInt(s, 10, 64)
					return i, nil
				}
				return nil, NotMatchErr
			},
		},
		{
			SrcType: Int64,
			DstType: String,
			Fn: func(src interface{}) (interface{}, error) {
				if s, ok := src.(int64); ok {
					return strconv.FormatInt(s, 64), nil
				}
				return nil, NotMatchErr
			},
		},
	}
)

// CopierWithOptions Set conversion rules： time-int64、time-string、int64-time、string-time
func CopierWithOptions(toValue interface{}, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    false,
		Converters:  tcs,
	})
}
