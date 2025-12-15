package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta        `path:"/api/v1/tunnel/create" method:"post" summary:"创建隧道"`
	Name          string  `json:"name" v:"required#隧道名称不能为空" dc:"隧道名称"`
	InNodeId      int     `json:"inNodeId" v:"required#入节点ID不能为空" dc:"入节点ID"`
	Type          int     `json:"type" v:"required#隧道类型不能为空" dc:"隧道类型"`
	OutNodeId     int     `json:"outNodeId" v:"required-unless:Type,1#出节点ID不能为空" dc:"出节点ID"`
	Flow          int     `json:"flow" v:"required#流量计算类型不能为空" dc:"流量计算类型"`
	TrafficRatio  float64 `json:"trafficRatio" v:"required#流量比例不能为空|between:0.0,100.0#流量比例必须在0.0-100.0之间" dc:"流量比例" d:"1.0"`
	InterfaceName string  `json:"interfaceName" dc:"接口名称"`
	Protocol      string  `json:"protocol" dc:"协议" d:"tls"`
	TcpListenAddr string  `json:"tcpListenAddr" dc:"TCP监听地址" d:"0.0.0.0"`
	UdpListenAddr string  `json:"udpListenAddr" dc:"UDP监听地址" d:"0.0.0.0"`
}

type CreateRes string
