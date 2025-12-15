package viteConfig

import (
	"context"
	"flux-panel-go/internal/dao"
	"flux-panel-go/internal/model/entity"
	"flux-panel-go/internal/model/localErr"

	"flux-panel-go/api/viteConfig/v1"
)

func (c *ControllerV1) GetViteConfigs(ctx context.Context, req *v1.GetViteConfigsReq) (res *v1.GetViteConfigsRes, err error) {
	viteConfigs := new([]*entity.ViteConfig)
	err = dao.ViteConfig.Ctx(ctx).Scan(viteConfigs)
	if err != nil {
		return nil, &localErr.CommonError{
			Code:   -1,
			ErrMsg: ERROR_GET_VITE_CONFIGS_FAILED,
		}
	}
	configMap := make(map[string]string)
	for _, config := range *viteConfigs {
		configMap[config.Name] = config.Value
	}
	temp := v1.GetViteConfigsRes(configMap)
	res = &temp
	return res, nil
}
