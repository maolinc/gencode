package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)

func (l *Create{{.CamelName}}Logic) Create{{.CamelName}}(req *types.Create{{.CamelName}}Req) (resp *types.Create{{.CamelName}}Resp, err error) {

	data := &model.{{.CamelName}}{}
	_ = copier.Copiers(data, req)
	err = l.svcCtx.{{.CamelName}}Model.Insert(l.ctx, nil, data)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Create{{.CamelName}} req: %v, error: %v", req, err)
	}

	return &types.Create{{.CamelName}}Resp{}, nil
}
