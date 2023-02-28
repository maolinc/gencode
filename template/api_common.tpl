
syntax = "{{.Syntax}}"

info(
	title: "view"
	desc: "抽取类型,方便在其他api使用"
	author: "{{.Author}}"
	email: "{{.Email}}"
	version: "{{.Version}}"
)

// 基本查询参数
type SearchBase {
    Id int64 `json:"id,optional"`
    Ids []int64 `json:"ids,optional"` // id集合
    UserId int64 `json:"userId,optional"`             // 用户id
    Keyword string `json:"keyword,optional"`          // 关键字
    PageSize int `json:"pageSize,optional"`           // 每页条数
    PageCurrent int `json:"pageCurrent,optional"`     // 当前页
    StartTime int64 `json:"startTime,optional"`       // 开始时间
    EndTime int64 `json:"endTime,optional"`           // 结束时间
    OrderSort []string `json:"orderSort,optional"`    // 排序
    SearchPlus [][]string `json:"searchPlus,optional"` // 加强版搜索参数  如 [["p_id", "=", "a", "string"], ["complete_time", ">=", "1674373544","number"]]
}

// 统一分页返回
type PageBase {
    Total int64 `json:"total,omitempty"` // 总条数
    PageCurrent int `json:"pageCurrent,omitempty"` // 当前页
    PageSize int `json:"pageSize,omitempty"` // 分页大小
    PageTotal int `json:"pageTotal,omitempty"` // 总分页数
}

type (
    IdReq {
        Id int64 `json:"id"`
    }
    IdsReq {
        Ids []int64 `json:"ids"`
    }
)

{{range  .Dataset.TableSet}}
type {{.CamelName}}View {
{{range  .Fields -}}{{if isIgnore 4 .IgnoreValue}}    {{.CamelName}} {{.DataType}} `json:"{{.StyleName}},optional"` //{{.Comment}}
{{end -}}
{{end -}}
}
{{end}}