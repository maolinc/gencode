package mexternalserver

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMExternalServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMExternalServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMExternalServerLogic {
	return &UpdateMExternalServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMExternalServerLogic) UpdateMExternalServer(req *types.UpdateMExternalServerReq) (resp *types.UpdateMExternalServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
