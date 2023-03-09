package martifact

import (
	"context"
	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMArtifactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMArtifactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMArtifactLogic {
	return &CreateMArtifactLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *CreateMArtifactLogic) CreateMArtifact(req *types.CreateMArtifactReq) (resp *types.CreateMArtifactResp, err error) {
	// todo: add your logic here and delete this line

	return &types.CreateMArtifactResp{}, nil
}
