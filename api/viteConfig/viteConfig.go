// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package viteConfig

import (
	"context"

	"flux-panel-go/api/viteConfig/v1"
)

type IViteConfigV1 interface {
	GetViteConfigs(ctx context.Context, req *v1.GetViteConfigsReq) (res *v1.GetViteConfigsRes, err error)
	GetViteConfig(ctx context.Context, req *v1.GetViteConfigReq) (res *v1.GetViteConfigRes, err error)
	UpdateViteConfigs(ctx context.Context, req *v1.UpdateViteConfigsReq) (res *v1.UpdateViteConfigsRes, err error)
	UpdateSingleViteConfig(ctx context.Context, req *v1.UpdateSingleViteConfigReq) (res *v1.UpdateSingleViteConfigRes, err error)
}
