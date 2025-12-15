package viteConfig

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/do"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/viteConfig/v1"
)

func (c *ControllerV1) GetViteConfig(ctx context.Context, req *v1.GetViteConfigReq) (res *v1.GetViteConfigRes, err error) {
	viteConfig := new(entity.ViteConfig)
	err = dao.ViteConfig.Ctx(ctx).Where(do.ViteConfig{
		Name: req.Name,
	}).Scan(viteConfig)
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorConfigNotFound,
		}
	}
	return &v1.GetViteConfigRes{
		Config: viteConfig,
	}, nil
}
