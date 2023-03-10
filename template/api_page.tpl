package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)


func (l *{{.CamelName}}PageLogic) {{.CamelName}}Page(req *types.Search{{.CamelName}}Req) (resp *types.Search{{.CamelName}}Resp, err error) {

	var cond model.{{.CamelName}}Query
	_ = copier.Copiers(&cond, req)

	list, err := l.svcCtx.{{.CamelName}}Model.FindListByPage(l.ctx, &cond)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "{{.CamelName}}Page req: %v, error: %v", req, err)
	}

	resp = &types.Search{{.CamelName}}Resp{}
	_ = copier.Copiers(&resp.List, list)

	return resp, nil
}