{{ $prefix := .Prefix }} {{ $serviceName := .ServiceName }}
syntax = "{{.Syntax}}"

info (
	title: "{{.CamelName}}"
	desc: "{{.Comment}}"
	author: "{{.Author}}"
	email: "{{.Email}}"
	version: "{{.Version}}"
)

import (
    "types/common.api"
)

//-----------------------{{.CamelName}}的接口-----------------------
@server (
	prefix: {{$prefix}}/{{.CamelName}}
	group: {{toLower .CamelName}}
)
service {{$serviceName}} {
	@doc "添加{{.CamelName}}"
	@handler Create{{.CamelName}}
	post /create (Create{{.CamelName}}Req) returns (Create{{.CamelName}}Resp)

    @doc "删除{{.CamelName}}"
    @handler Delete{{.CamelName}}
    post /delete (Delete{{.CamelName}}Req) returns (Delete{{.CamelName}}Resp)

    @doc "查询{{.CamelName}}详情"
    @handler Detail{{.CamelName}}
    post /detail (Detail{{.CamelName}}Req) returns (Detail{{.CamelName}}Resp)

	@doc "分页查询{{.CamelName}}"
	@handler Page{{.CamelName}}
	post /page (Search{{.CamelName}}Req) returns (Search{{.CamelName}}Resp)

	@doc "更新{{.CamelName}}"
	@handler Update{{.CamelName}}
	post /update (Update{{.CamelName}}Req) returns (Update{{.CamelName}}Resp)
}


//-----------------------请求、响应数据-----------------------

type {{.CamelName}}View {
{{range  .Fields -}}{{if isIgnore 4 .IgnoreValue}}    {{.CamelName}} {{.DataType}} `json:"{{.StyleName}}"` //{{.Comment}}
{{end -}}
{{end -}}
}


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