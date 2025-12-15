package viteConfig

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/viteConfig/v1"

	"github.com/gogf/gf/v2/database/gdb"
)

func (c *ControllerV1) UpdateViteConfigs(ctx context.Context, req *v1.UpdateViteConfigsReq) (res *v1.UpdateViteConfigsRes, err error) {
	if req.Configs == nil || len(req.Configs) == 0 {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorConfigsRequired,
		}
	}
	err = dao.ViteConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for key, value := range req.Configs {
			if key == "" {
				continue
			}
			_, sqlErr := dao.ViteConfig.Ctx(ctx).Where(dao.ViteConfig.Columns().Name, key).Update(dao.ViteConfig.Columns().Value, value)
			if sqlErr != nil {
				return sqlErr
			}
		}
		return nil
	})
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ErrorUpdateViteConfigsFailed,
		}
	}
	r := v1.UpdateViteConfigsRes(SuccessUpdateViteConfigs)
	return &r, nil
}
