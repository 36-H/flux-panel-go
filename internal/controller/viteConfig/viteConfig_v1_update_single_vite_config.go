package viteConfig

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/viteConfig/v1"
)

func (c *ControllerV1) UpdateSingleViteConfig(ctx context.Context, req *v1.UpdateSingleViteConfigReq) (res *v1.UpdateSingleViteConfigRes, err error) {
	if req.Name == "" {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorNameRequired,
		}
	}
	if req.Value == "" {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorValueRequired,
		}
	}
	_, err = dao.ViteConfig.Ctx(ctx).Where(dao.ViteConfig.Columns().Name, req.Name).Update(dao.ViteConfig.Columns().Value, req.Value)
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorUpdateViteConfigsFailed,
		}
	}
	r := v1.UpdateSingleViteConfigRes(SuccessUpdateViteConfigs)
	return &r, nil
}
