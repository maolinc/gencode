package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)

func (l *Create{{.CamelName}}Logic) Create{{.CamelName}}(in *pb.Create{{.CamelName}}Req) (*pb.Create{{.CamelName}}Resp, error) {

	data := &model.{{.CamelName}}{}
	_ = copier.Copiers(data, in)
	err := l.svcCtx.{{.CamelName}}Model.Insert(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Create{{.CamelName}} req: %v, error: %v", in, err)
	}

	return &pb.Create{{.CamelName}}Resp{}, nil
}

