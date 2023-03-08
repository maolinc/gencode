package mexternaldiffrecords

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMExternalDiffRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMExternalDiffRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMExternalDiffRecordsLogic {
	return &UpdateMExternalDiffRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMExternalDiffRecordsLogic) UpdateMExternalDiffRecords(req *types.UpdateMExternalDiffRecordsReq) (resp *types.UpdateMExternalDiffRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
