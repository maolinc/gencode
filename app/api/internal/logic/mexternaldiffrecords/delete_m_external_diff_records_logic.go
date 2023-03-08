package mexternaldiffrecords

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMExternalDiffRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMExternalDiffRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMExternalDiffRecordsLogic {
	return &DeleteMExternalDiffRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMExternalDiffRecordsLogic) DeleteMExternalDiffRecords(req *types.IdsReq) (resp *types.DeleteMExternalDiffRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
