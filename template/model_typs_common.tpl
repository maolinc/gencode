package model

type SearchItem struct {
    Table    string `json:"table"`
	Field    string `json:"field"`          // 字段
	Value    string `json:"value"`          // 值
	Type     string `json:"type"`           // 值的数据类型 number string date numberArray stringArray
	Operator string `json:"operator"`       // 操作符 = != > >= 包含 不包含...
	Logic    string `json:"logic,optional"` // 逻辑符 and | or
}

type SearchGroup struct {
	Group []SearchItem `json:"group"`          // 条件组合
	Logic string       `json:"logic,optional"` // 逻辑符 and | or
}

type SearchBase struct {
	Keyword     string        `json:"keyword,optional"`               // 关键字
	Cursor      int64         `json:"cursor,optional"`                // 游标,基于游标分页时使用
	CursorAsc   bool          `json:"cursorAsc,optional"`             // 游标分页时方向 true:asc  false:desc
	PageSize    int           `json:"pageSize,default=20,optional"`   // 每页条数
	PageCurrent int           `json:"pageCurrent,default=1,optional"` // 当前页
	OrderSort   []string      `json:"orderSort,optional"`             // 排序 eg： ["create_time asc", "id desc"]
	SearchPlus  []SearchGroup `json:"searchPlus,optional"`            // 加强版自定义搜索参数
}