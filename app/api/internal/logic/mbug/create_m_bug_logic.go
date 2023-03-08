package mbug

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMBugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMBugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMBugLogic {
	return &CreateMBugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMBugLogic) CreateMBug(req *types.CreateMBugReq) (resp *types.CreateMBugResp, err error) {
	// todo: add your logic here and delete this line

	return
}
