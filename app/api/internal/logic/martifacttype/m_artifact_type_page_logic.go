package martifacttype

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MArtifactTypePageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMArtifactTypePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MArtifactTypePageLogic {
	return &MArtifactTypePageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MArtifactTypePageLogic) MArtifactTypePage(req *types.SearchMArtifactTypeReq) (resp *types.SearchMArtifactTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
