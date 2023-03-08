package mexternalserver

import (
	"context"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MExternalServerPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMExternalServerPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MExternalServerPageLogic {
	return &MExternalServerPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MExternalServerPageLogic) MExternalServerPage(req *types.SearchMExternalServerReq) (resp *types.SearchMExternalServerResp, err error) {
	// todo: add your logic here and delete this line

	return
}
