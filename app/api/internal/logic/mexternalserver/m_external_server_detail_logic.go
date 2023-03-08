package mexternalserver

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MExternalServerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMExternalServerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MExternalServerDetailLogic {
	return &MExternalServerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MExternalServerDetailLogic) MExternalServerDetail(req *types.IdReq) (resp *types.DetailMExternalServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
