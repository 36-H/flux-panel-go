package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"flux-panel-go/internal/controller/node"
	"flux-panel-go/internal/controller/user"
	"flux-panel-go/internal/logic/middleware"
	"flux-panel-go/internal/utils"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// Initialize JWT utility
			jwtSecret := g.Cfg().MustGet(ctx, "jwt.secret").String()
			if jwtSecret == "" {
				jwtSecret = "flux-panel-secret-key"
			}
			utils.InitJwtUtil(jwtSecret)
			// 增加对数据库的连接检测
			if _, err = g.DB().Query(ctx, "SELECT 1"); err != nil {
				g.Log().Fatal(ctx, "数据库连接失败:", err)
			}
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.ResponseHandler)
				group.Bind(
					user.NewV1(),
					node.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
