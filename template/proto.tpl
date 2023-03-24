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

message SearchItem {
    string field=1; // 字段
    string value=2; // 值
    string type=3; // 值的数据类型 number string date numberArray stringArray
    string operator=4; // 操作符 = != > >= 包含 不包含...
    string logic=5; // 逻辑符 and | or
    string table=6; // 表
}
message SearchGroup {
    repeated SearchItem group=1; // 条件组合
    string logic=2; // 逻辑符 and | or
}


{{range .Dataset.TableSet}}
//-----------------------{{.Name}}-----------------------
message {{.CamelName}}View {
{{range $index,$value := .Fields}}{{if isIgnore 4 .IgnoreValue}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Create{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if isIgnore 1 .IgnoreValue}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Create{{.CamelName}}Resp {
}

message Update{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if isIgnore 2 .IgnoreValue}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
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
    int64 cursor = 1; // 分页游标
    bool cursorAsc = 2; // 游标分页时方向 true:asc  false:desc
    int64 pageSize = 3; // 每页条数
    int64 pageCurrent = 4;  // 当前页
    repeated string orderSort = 5;  // 排序 eg： ["create_time asc", "id desc"]
    repeated SearchGroup searchPlus = 6; // 加强版搜索参数
{{range $index,$value := .Fields}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 8}}; //{{.Comment}}
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