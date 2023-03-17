package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)

func (l *Page{{.CamelName}}Logic) Page{{.CamelName}}(in *pb.Search{{.CamelName}}Req) (*pb.Search{{.CamelName}}Resp, error) {

	var cond model.{{.CamelName}}Query
	_ = copier.Copiers(&cond, in)

	list, err := l.svcCtx.{{.CamelName}}Model.FindListByPage(l.ctx, &cond)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Page{{.CamelName}} req: %v, error: %v", in, err)
	}

	resp := &pb.Search{{.CamelName}}Resp{}
	_ = copier.Copiers(&resp.List, list)

	return resp, nil
}
