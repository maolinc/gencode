package martifacttype

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMArtifactTypeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMArtifactTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMArtifactTypeLogic {
	return &CreateMArtifactTypeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMArtifactTypeLogic) CreateMArtifactType(req *types.CreateMArtifactTypeReq) (resp *types.CreateMArtifactTypeResp, err error) {
	// todo: add your logic here and delete this line

	return
}
