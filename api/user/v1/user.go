package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta   `path:"/api/v1/user/login" tags:"Users" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

type LoginRes struct {
	Token                 string `json:"token" dc:"token"`
	Name                  string `json:"name" dc:"用户名"`
	RoleId                int    `json:"role_id" dc:"角色ID"`
	RequirePasswordChange bool   `json:"require_password_change" dc:"是否需要修改密码"`
}
