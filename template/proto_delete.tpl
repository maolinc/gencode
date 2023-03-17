package {{.Package}}

import (
	"github.com/pkg/errors"
)


func (l *Delete{{.CamelName}}Logic) Delete{{.CamelName}}(in *pb.Delete{{.CamelName}}Req) (*pb.Delete{{.CamelName}}Resp, error) {

	err := l.svcCtx.{{.CamelName}}Model.Delete(l.ctx, {{.PrimaryFmtV}})
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "Delete{{.CamelName}} req: %v, error: %v", in, err)
	}
	
	return &pb.Delete{{.CamelName}}Resp{}, nil
}
