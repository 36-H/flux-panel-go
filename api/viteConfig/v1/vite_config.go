package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetViteConfigsReq struct {
	g.Meta `path:"/api/v1/config/list" tags:"Vite Config" method:"get" summary:"获取所有网站配置"`
}

type GetViteConfigsRes map[string]string
