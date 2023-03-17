package {{.Package}}

import (
	"github.com/maolinc/copier"
	"github.com/pkg/errors"
)


func (l *Detail{{.CamelName}}Logic) Detail{{.CamelName}}(in *pb.Detail{{.CamelName}}Req) (*pb.Detail{{.CamelName}}Resp, error) {

	data, err := l.svcCtx.{{.CamelName}}Model.FindOne(l.ctx, {{.PrimaryFmtV}})
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail"), "{{.CamelName}}Detail req: %v, error: %v", in, err)
	}
	if data == nil {
		return nil, errors.New("resource not exist")
	}

	resp := &pb.Detail{{.CamelName}}Resp{}
	_ = copier.Copiers(&resp, data)

	return resp, nil
}
