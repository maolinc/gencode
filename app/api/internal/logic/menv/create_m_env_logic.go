package menv

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/maolinc/gencode/app/model"
	"github.com/pkg/errors"

	"github.com/maolinc/gencode/app/api/internal/svc"
	"github.com/maolinc/gencode/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMEnvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMEnvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMEnvLogic {
	return &CreateMEnvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMEnvLogic) CreateMEnv(req *types.CreateMEnvReq) (resp *types.CreateMEnvResp, err error) {
	// todo: add your logic here and delete this line

	data := &model.MEnv{}
	_ = copier.Copy(data, req)
	err = l.svcCtx.MEnvModel.Insert(l.ctx, nil, data)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %v", req)
	}

	return &types.CreateMEnvResp{}, nil
}
