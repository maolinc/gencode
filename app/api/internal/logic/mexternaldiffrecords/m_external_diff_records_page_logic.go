package mexternaldiffrecords

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MExternalDiffRecordsPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMExternalDiffRecordsPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MExternalDiffRecordsPageLogic {
	return &MExternalDiffRecordsPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MExternalDiffRecordsPageLogic) MExternalDiffRecordsPage(req *types.SearchMExternalDiffRecordsReq) (resp *types.SearchMExternalDiffRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
