package mexternaldiffrecords

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMExternalDiffRecordsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMExternalDiffRecordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMExternalDiffRecordsLogic {
	return &CreateMExternalDiffRecordsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMExternalDiffRecordsLogic) CreateMExternalDiffRecords(req *types.CreateMExternalDiffRecordsReq) (resp *types.CreateMExternalDiffRecordsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
