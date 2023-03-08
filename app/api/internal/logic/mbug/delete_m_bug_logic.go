package mbug

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMBugLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMBugLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMBugLogic {
	return &DeleteMBugLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMBugLogic) DeleteMBug(req *types.IdsReq) (resp *types.DeleteMBugResp, err error) {
	// todo: add your logic here and delete this line

	return
}
