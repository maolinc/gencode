package mexternaltask

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MExternalTaskPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMExternalTaskPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MExternalTaskPageLogic {
	return &MExternalTaskPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MExternalTaskPageLogic) MExternalTaskPage(req *types.SearchMExternalTaskReq) (resp *types.SearchMExternalTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
