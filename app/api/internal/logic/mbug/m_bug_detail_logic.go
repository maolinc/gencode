package mbug

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MBugDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMBugDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MBugDetailLogic {
	return &MBugDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MBugDetailLogic) MBugDetail(req *types.IdReq) (resp *types.DetailMBugResp, err error) {
	// todo: add your logic here and delete this line

	return
}
