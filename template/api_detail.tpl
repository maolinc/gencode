package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)

func (l *{{.CamelName}}DetailLogic) {{.CamelName}}Detail(req *types.IdReq) (resp *types.Detail{{.CamelName}}Resp, err error) {

	data, err := l.svcCtx.{{.CamelName}}Model.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "{{.CamelName}}Detail req: %v, error: %v", req, err)
	}
	if data == nil {
		return nil, errors.New("resource not exist")
	}

	resp = &types.Detail{{.CamelName}}Resp{}
	_ = copier.Copiers(&resp, data)

	return resp, nil
}
