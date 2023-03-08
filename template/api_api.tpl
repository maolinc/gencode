{{ $prefix := .Prefix }} {{ $serviceName := .ServiceName }}
syntax = "{{.Syntax}}"

info(
	title: "{{$serviceName}}"
	desc: "api接口文件"
	author: "{{.Author}}"
	email: "{{.Email}}"
	version: "{{.Version}}"
)

import (
    "types/common.api"
{{range .Dataset.TableSet}}    "types/{{toLower .CamelName}}.api"
{{end}})

{{range .Dataset.TableSet}}
//-----------------------{{.CamelName}}的接口-----------------------
@server(
	prefix: {{$prefix}}/{{.CamelName}}
	group: {{toLower .CamelName}}
)
service {{$serviceName}} {
	@doc "添加{{.CamelName}}"
	@handler Create{{.CamelName}}
	post /create (Create{{.CamelName}}Req) returns (Create{{.CamelName}}Resp)

    @doc "删除{{.CamelName}}"
    @handler Delete{{.CamelName}}
    post /delete (IdsReq) returns (Delete{{.CamelName}}Resp)

    @doc "查询{{.CamelName}}详情"
    @handler {{.CamelName}}Detail
    post /detail (IdReq) returns (Detail{{.CamelName}}Resp)

	@doc "分页查询{{.CamelName}}"
	@handler {{.CamelName}}Page
	post /page (Search{{.CamelName}}Req) returns (Search{{.CamelName}}Resp)

	@doc "更新{{.CamelName}}"
	@handler Update{{.CamelName}}
	post /update (Update{{.CamelName}}Req) returns (Update{{.CamelName}}Resp)
}
{{end}}