{{ $prefix := .Prefix }}
syntax = "{{.Syntax}}"

info(
	title: "{{.CamelName}}类型"
	desc: "{{.Comment}}的类型"
	author: "{{.Author}}"
	email: "{{.Email}}"
	version: "{{.Version}}"
)

import (
    "types/common.api"
)

type Create{{.CamelName}}Req {
{{- range  .Fields -}}{{if isIgnore 1 .IgnoreValue}}    {{if not .IsPrimary}} {{.CamelName}} {{.DataType}} `json:"{{.StyleName}}{{- if .HasDefault -}},default={{.Default}}{{end}}{{- if .IsNullable -}},optional{{end}}"` //{{.Comment}} {{end}}
{{end -}}
{{end -}}}

type Create{{.CamelName}}Resp {
}

type Update{{.CamelName}}Req {
{{range  .Fields -}}{{if isIgnore 2 .IgnoreValue}}    {{.CamelName}} {{.DataType}} `json:"{{.StyleName}}{{- if not .IsPrimary -}},optional{{end}}"` //{{.Comment}}
{{end -}}
{{end -}}}

type Update{{.CamelName}}Resp {
}


type Delete{{.CamelName}}Req {
{{range  .Fields -}}{{if .IsPrimary}}    {{.CamelName}} {{.DataType}} `json:"{{.StyleName}}"` //{{.Comment}}
{{end -}}
{{end -}}}

type Delete{{.CamelName}}Resp {
}

type Detail{{.CamelName}}Req {
{{range  .Fields -}}{{if .IsPrimary}}    {{.CamelName}} {{.DataType}} `json:"{{.StyleName}}"` //{{.Comment}}
{{end -}}
{{end -}}}

type Detail{{.CamelName}}Resp {
    {{.CamelName}}View
}

type Search{{.CamelName}}Req {
    SearchBase // 内置查询参数
    {{range .Fields -}}  {{.CamelName}} {{.DataType}} `json:"{{.StyleName}},optional"` //{{.Comment}}
    {{end -}}
}

type Search{{.CamelName}}Resp {
    PageBase // 分页参数
    List []{{.CamelName}}View `json:"list"` // 列表
}