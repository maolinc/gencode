package mbug

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MBugPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMBugPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MBugPageLogic {
	return &MBugPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MBugPageLogic) MBugPage(req *types.SearchMBugReq) (resp *types.SearchMBugResp, err error) {
	// todo: add your logic here and delete this line

	return
}
