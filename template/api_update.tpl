package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)

func (l *Update{{.CamelName}}Logic) Update{{.CamelName}}(req *types.Update{{.CamelName}}Req) (resp *types.Update{{.CamelName}}Resp, err error) {

	var data model.{{.CamelName}}
	_ = copier.Copiers(&data, req)

	err = l.svcCtx.{{.CamelName}}Model.Update(l.ctx, nil, &data)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Update{{.CamelName}} req: %v, error: %v", req, err)
	}
	
	return &types.Update{{.CamelName}}Resp{}, nil
}
