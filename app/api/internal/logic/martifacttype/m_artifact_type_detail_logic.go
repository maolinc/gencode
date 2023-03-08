package martifacttype

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MArtifactTypeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMArtifactTypeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MArtifactTypeDetailLogic {
	return &MArtifactTypeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MArtifactTypeDetailLogic) MArtifactTypeDetail(req *types.IdReq) (resp *types.DetailMArtifactTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
