package mbug

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMBugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMBugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMBugLogic {
	return &UpdateMBugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMBugLogic) UpdateMBug(req *types.UpdateMBugReq) (resp *types.UpdateMBugResp, err error) {
	// todo: add your logic here and delete this line

	return
}
