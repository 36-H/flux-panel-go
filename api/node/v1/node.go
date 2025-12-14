package v1

import (
	"flux-panel-go/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta    `path:"/api/v1/node/create" tags:"Nodes" method:"post" summary:"创建节点"`
	Name      string `json:"name" v:"required#节点名称不能为空" dc:"节点名称"`
	IP        string `json:"ip" v:"required#入口IP不能为空|ip#入口IP格式错误" dc:"入口IP"`
	ServerIp  string `json:"server_ip" v:"required#服务器IP不能为空|ip#服务器IP格式错误" dc:"服务器IP"`
	PortStart int    `json:"port_start" v:"required#起始端口不能为空|min:1#起始端口必须大于0|max:65532#起始端口必须小于65532" dc:"起始端口"`
	PortEnd   int    `json:"port_end" v:"required#结束端口不能为空|min:1#结束端口必须大于0|max:65532#结束端口必须小于65532" dc:"结束端口"`
}

type CreateRes string

type GetListReq struct {
	g.Meta `path:"/api/v1/node/list" tags:"Nodes" method:"get" summary:"获取节点列表"`
}

type GetListRes struct {
	Nodes []*entity.Node `json:"nodes"`
}
