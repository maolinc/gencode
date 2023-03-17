
{{range  .Dataset.TableSet}}
type {{.CamelName}}Query struct {
    SearchBase
    *{{.CamelName}}
}
{{end}}