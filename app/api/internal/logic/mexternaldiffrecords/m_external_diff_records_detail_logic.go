package mexternaldiffrecords

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MExternalDiffRecordsDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMExternalDiffRecordsDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MExternalDiffRecordsDetailLogic {
	return &MExternalDiffRecordsDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MExternalDiffRecordsDetailLogic) MExternalDiffRecordsDetail(req *types.IdReq) (resp *types.DetailMExternalDiffRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
