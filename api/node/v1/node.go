package v1

import (
	"flux-panel-go/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"/api/v1/node/create" tags:"Nodes" method:"post" summary:"创建节点"`
	Name     string `json:"name" v:"required#节点名称不能为空" dc:"节点名称"`
	IP       string `json:"ip" v:"required#入口IP不能为空|ip#入口IP格式错误" dc:"入口IP"`
	ServerIp string `json:"server_ip" v:"required#服务器IP不能为空|ip#服务器IP格式错误" dc:"服务器IP"`
	PortSta  int    `json:"port_sta" v:"required#起始端口不能为空|min:1#起始端口必须大于0|max:65532#起始端口必须小于65532" dc:"起始端口"`
	PortEnd  int    `json:"port_end" v:"required#结束端口不能为空|min:1#结束端口必须大于0|max:65532#结束端口必须小于65532" dc:"结束端口"`
}

type CreateRes string

type GetListReq struct {
	g.Meta `path:"/api/v1/node/list" tags:"Nodes" method:"get" summary:"获取节点列表"`
}

type GetListRes struct {
	Nodes []*entity.Node `json:"nodes"`
}

type UpdateReq struct {
	g.Meta   `path:"/api/v1/node/update" tags:"Nodes" method:"put" summary:"更新节点"`
	ID       int    `json:"id" v:"required#节点ID不能为空" dc:"节点ID"`
	Name     string `json:"name" v:"required#节点名称不能为空" dc:"节点名称"`
	IP       string `json:"ip" v:"required#入口IP不能为空|ip#入口IP格式错误" dc:"入口IP"`
	ServerIp string `json:"server_ip" v:"required#服务器IP不能为空|ip#服务器IP格式错误" dc:"服务器IP"`
	PortSta  int    `json:"port_sta" v:"required#起始端口不能为空|min:1#起始端口必须大于0|max:65532#起始端口必须小于65532" dc:"起始端口"`
	PortEnd  int    `json:"port_end" v:"required#结束端口不能为空|min:1#结束端口必须大于0|max:65532#结束端口必须小于65532" dc:"结束端口"`
	Http     *int   `json:"http"  dc:"HTTP端口"`
	Tls      *int   `json:"tls"  dc:"TLS端口"`
	Socks    *int   `json:"socks"  dc:"SOCKS端口"`
}

type UpdateRes string

type DeleteReq struct {
	g.Meta `path:"/api/v1/node/delete" tags:"Nodes" method:"delete" summary:"删除节点"`
	ID     int `json:"id" v:"required#节点ID不能为空" dc:"节点ID"`
}

type DeleteRes string

type GetInstallCommandReq struct {
	g.Meta `path:"/api/v1/node/install-command" tags:"Nodes" method:"get" summary:"获取节点安装命令"`
	ID     int `json:"id" v:"required#节点ID不能为空" dc:"节点ID"`
}
type GetInstallCommandRes string
