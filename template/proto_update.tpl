package {{.Package}}

import (
	"github.com/maolinc/copier"
	{{.ModelPkg}}
	"github.com/pkg/errors"
)

func (l *Update{{.CamelName}}Logic) Update{{.CamelName}}(in *pb.Update{{.CamelName}}Req) (*pb.Update{{.CamelName}}Resp, error) {

	var data model.{{.CamelName}}
	_ = copier.Copiers(&data, in)

	err := l.svcCtx.{{.CamelName}}Model.Update(l.ctx, &data)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Update{{.CamelName}} req: %v, error: %v", in, err)
	}
	
	return &pb.Update{{.CamelName}}Resp{}, nil
}
