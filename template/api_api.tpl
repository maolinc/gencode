{{ $prefix := .Prefix }} {{ $serviceName := .ServiceName }}
syntax = "{{.Syntax}}"

info(
	title: "{{$serviceName}}"
	desc: "api入口文件"
	author: "{{.Author}}"
	email: "{{.Email}}"
	version: "{{.Version}}"
)

import (
{{range .Dataset.TableSet}}    "types/{{toLower .CamelName}}.api"
{{end}})