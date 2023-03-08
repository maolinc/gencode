package martifacttype

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMArtifactTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMArtifactTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMArtifactTypeLogic {
	return &DeleteMArtifactTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMArtifactTypeLogic) DeleteMArtifactType(req *types.IdsReq) (resp *types.DeleteMArtifactTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
