package {{.Package}}

import (
	"github.com/pkg/errors"
)


func (l *Delete{{.CamelName}}Logic) Delete{{.CamelName}}(req *types.Delete{{.CamelName}}Req) (resp *types.Delete{{.CamelName}}Resp, err error) {

	err = l.svcCtx.{{.CamelName}}Model.Delete(l.ctx, {{.PrimaryFmtV}})
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Delete{{.CamelName}} req: %v, error: %v", req, err)
	}
	
	return &types.Delete{{.CamelName}}Resp{}, nil
}
