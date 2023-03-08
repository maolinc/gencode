package mexternaltask

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MExternalTaskDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMExternalTaskDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MExternalTaskDetailLogic {
	return &MExternalTaskDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MExternalTaskDetailLogic) MExternalTaskDetail(req *types.IdReq) (resp *types.DetailMExternalTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
