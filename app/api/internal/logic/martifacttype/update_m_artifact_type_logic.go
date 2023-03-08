package martifacttype

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMArtifactTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMArtifactTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMArtifactTypeLogic {
	return &UpdateMArtifactTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMArtifactTypeLogic) UpdateMArtifactType(req *types.UpdateMArtifactTypeReq) (resp *types.UpdateMArtifactTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
