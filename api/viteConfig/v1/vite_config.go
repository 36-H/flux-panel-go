package v1

import (
	"flux-panel-go/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetViteConfigsReq struct {
	g.Meta `path:"/api/v1/config/list" tags:"Vite Config" method:"get" summary:"获取所有网站配置"`
}

type GetViteConfigsRes map[string]string

type GetViteConfigReq struct {
	g.Meta `path:"/api/v1/config" tags:"Vite Config" method:"get" summary:"获取指定网站配置"`
	Name   string `json:"name" v:"required#请输入配置名称" dc:"配置名称"`
}

type GetViteConfigRes struct {
	Config *entity.ViteConfig `json:"config"`
}

type UpdateViteConfigsReq struct {
	g.Meta  `path:"/api/v1/config" tags:"Vite Config" method:"put" summary:"更新网站配置"`
	Configs map[string]string `json:"configs" v:"required#请输入配置" dc:"配置"`
}

type UpdateViteConfigsRes string
