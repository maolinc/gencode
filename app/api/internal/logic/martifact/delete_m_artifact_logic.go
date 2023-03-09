package martifact

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMArtifactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMArtifactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMArtifactLogic {
	return &DeleteMArtifactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMArtifactLogic) DeleteMArtifact(req *types.IdsReq) (resp *types.DeleteMArtifactResp, err error) {
	// todo: add your logic here and delete this line

	return
}
