func (l *CreateMArtifactLogic) CreateMArtifact(req *types.CreateMArtifactReq) (resp *types.CreateMArtifactResp, err error) {

	var data model.MArtifact
	_ = copier.Copiers(&data, req)

	err = l.svcCtx.MArtifactModel.Insert(l.ctx, nil, &data)
	if err != nil {
		return nil, errors.Wrapf(errors.New("operate fail!"), "createMArtifact error req: %v, error: %v", req, err)
	}

	return &types.CreateMArtifactResp{}, nil
}
