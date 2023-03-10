syntax = "{{.Syntax}}";

option go_package ="{{.GoPackage}}";

package {{.Package}};

// ------------------------------------
// Rpc service
// ------------------------------------
{{range .Dataset.TableSet}}
//-----------------------{{.Name}}-----------------------
service {{.CamelName}}{
    rpc Create{{.CamelName}}(Create{{.CamelName}}Req) returns (Create{{.CamelName}}Resp);
    rpc Update{{.CamelName}}(Update{{.CamelName}}Req) returns (Update{{.CamelName}}Resp);
    rpc Delete{{.CamelName}}(Delete{{.CamelName}}Req) returns (Delete{{.CamelName}}Resp);
    rpc Detail{{.CamelName}}(Detail{{.CamelName}}Req) returns (Detail{{.CamelName}}Resp);
    rpc Page{{.CamelName}}(Search{{.CamelName}}Req) returns (Search{{.CamelName}}Resp);
}
{{end}}


// ------------------------------------
// Rpc message
// ------------------------------------

//-----------------------通用message-----------------------
message IdReq {
    int64 id = 1; //id
}

message IdsReq {
    int64 ids = 1; //ids
}

message PlusItem {
    repeated string item = 1;
}

message SearchBase {
    optional string keyword = 1; // 关键字
    optional int64 cursor = 2; // 分页游标
    optional bool cursorAsc = 3; // 游标分页时方向 true:asc  false:desc
    optional int64 pageSize = 4; // 每页条数
    optional int64 pageCurrent = 5;  // 当前页
    repeated string orderSort = 6;  // 排序 eg： ["create_time asc", "id desc"]
    repeated PlusItem searchPlus = 7; // 加强版搜索参数  eg: [["p_id", "=", "a", "string"], ["complete_time", ">=", "1674373544","number"]]
}


{{range .Dataset.TableSet}}
//-----------------------{{.Name}}-----------------------
message {{.CamelName}}View {
{{range $index,$value := .Fields}}{{if isIgnore 4 .IgnoreValue}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Create{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if isIgnore 1 .IgnoreValue}}    {{if .IsNullable -}}optional {{end}}{{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Create{{.CamelName}}Resp {
}

message Update{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if isIgnore 2 .IgnoreValue}}    {{if .IsNullable -}}optional {{end}}{{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Update{{.CamelName}}Resp {
}

message Delete{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if .IsPrimary}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Delete{{.CamelName}}Resp {
}

message Detail{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if .IsPrimary}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Detail{{.CamelName}}Resp {
{{range $index,$value := .Fields}}{{if isIgnore 4 .IgnoreValue}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Search{{.CamelName}}Req {
    SearchBase baseCond = 1; // 基本查询参数
{{range $index,$value := .Fields}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 2}}; //{{.Comment}}
{{end}}}

message Search{{.CamelName}}Resp {
    int64 total = 1;
    int64 pageCurrent = 2;
    int64 pageSize = 3;
    int64 pageTotal = 4;
    int64 lastCursor = 5;
    repeated {{.CamelName}}View list = 7; // 列表
}
{{end}}