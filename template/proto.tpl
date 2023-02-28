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
    rpc Delete{{.CamelName}}(IdsReq) returns (Delete{{.CamelName}}Resp);
    rpc Detail{{.CamelName}}(IdReq) returns (Detail{{.CamelName}}Resp);
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
    optional int64 id = 1;
    repeated int64 ids = 2;
    optional int64 userId = 3; // 用户id
    optional string keyword = 4; // 关键字
    optional int64 pageSize = 5; // 每页条数
    optional int64 pageCurrent = 6;  // 当前页
    optional int64 startTime = 7; // 开始时间
    optional int64 endTime = 8;  // 结束时间
    repeated string orderSort = 9;  // 排序
    repeated PlusItem searchPlus = 10; // 加强版参数
}


{{range .Dataset.TableSet}}
//-----------------------{{.Name}}-----------------------
message {{.CamelName}}View {
{{range $index,$value := .Fields}}{{if isIgnore 4 .IgnoreValue}}   {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Create{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if isIgnore 1 .IgnoreValue}}   {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Create{{.CamelName}}Resp {
}

message Update{{.CamelName}}Req {
{{range $index,$value := .Fields}}{{if isIgnore 2 .IgnoreValue}}   {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Update{{.CamelName}}Resp {
}

message Delete{{.CamelName}}Resp {
}

message Detail{{.CamelName}}Resp {
{{range $index,$value := .Fields}}{{if isIgnore 4 .IgnoreValue}}    {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 1}}; //{{.Comment}}
{{end -}}
{{end -}}}

message Search{{.CamelName}}Req {
    SearchBase baseCond = 1; // 基本查询参数
{{range $index,$value := .Fields}}   {{.DataType}} {{toCamelWithStartLower .CamelName}} = {{add $index 2}}; //{{.Comment}}
{{end}}}

message Search{{.CamelName}}Resp {
    int64 total = 1;
    int64 pageCurrent = 2;
    int64 pageSize = 3;
    int64 pageTotal = 4;
    repeated {{.CamelName}}View list = 7; // 列表
}
{{end}}