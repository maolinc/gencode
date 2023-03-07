package model

type SearchBase struct {
	Id int64 `json:","`
	// 游标
	Cursor    int64
	CursorAsc bool `json:",default=true"`
	// 每页条数
	PageSize int `json:",default=20"`
	// 当前页
	PageCurrent int `json:",default=1"`
	// 排序 eg： ["create_time asc", "id desc"]
	OrderSort []string `json:","`
	// 开始时间
	StartTime int64 `json:","`
	// 结束时间
	EndTime int64 `json:","`
	// plusCond  Benefits: It is not necessary to define each field separately, and the query conditions are flexible
	// [[field,  symbol, value, dataType]]，symbol:= != > >= < <= in like..., dataType:dataType of the value
	//[["name", "=", "服务器程序", "string"], ["complete_time", ">=", 1674373544, "date"],["id", "in", "1,2,3", "numberArray"]]
	SearchPlus [][]string `json:","`
}


{{range  .Dataset.TableSet}}
type {{.CamelName}}Query struct {
    SearchBase
    {{.CamelName}} *{{.CamelName}}
}
{{end}}