package martifact

import (
	"context"
	"github.com/maolinc/copier"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type MArtifactDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMArtifactDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MArtifactDetailLogic {
	return &MArtifactDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MArtifactDetailLogic) MArtifactDetail(req *types.IdReq) (resp *types.DetailMArtifactResp, err error) {

	data, err := l.svcCtx.MArtifactModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %v", req)
	}
	if data == nil {
		return nil, errors.New("resource not exist")
	}

	resp = &types.DetailMArtifactResp{}
	_ = copier.Copiers(&resp, data)

	return resp, nil
}
