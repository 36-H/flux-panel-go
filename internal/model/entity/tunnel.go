// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Tunnel is the golang structure for table tunnel.
type Tunnel struct {
	Id            int     `json:"id"            orm:"id"              description:""` //
	Name          string  `json:"name"          orm:"name"            description:""` //
	TrafficRatio  float64 `json:"trafficRatio"  orm:"traffic_ratio"   description:""` //
	InNodeId      int     `json:"inNodeId"      orm:"in_node_id"      description:""` //
	InIp          string  `json:"inIp"          orm:"in_ip"           description:""` //
	OutNodeId     int     `json:"outNodeId"     orm:"out_node_id"     description:""` //
	OutIp         string  `json:"outIp"         orm:"out_ip"          description:""` //
	Type          int     `json:"type"          orm:"type"            description:""` //
	Protocol      string  `json:"protocol"      orm:"protocol"        description:""` //
	Flow          int     `json:"flow"          orm:"flow"            description:""` //
	TcpListenAddr string  `json:"tcpListenAddr" orm:"tcp_listen_addr" description:""` //
	UdpListenAddr string  `json:"udpListenAddr" orm:"udp_listen_addr" description:""` //
	InterfaceName string  `json:"interfaceName" orm:"interface_name"  description:""` //
	CreatedTime   int64   `json:"createdTime"   orm:"created_time"    description:""` //
	UpdatedTime   int64   `json:"updatedTime"   orm:"updated_time"    description:""` //
	Status        int     `json:"status"        orm:"status"          description:""` //
}
