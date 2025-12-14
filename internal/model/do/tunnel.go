// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Tunnel is the golang structure of table tunnel for DAO operations like Where/Data.
type Tunnel struct {
	g.Meta        `orm:"table:tunnel, do:true"`
	Id            any //
	Name          any //
	TrafficRatio  any //
	InNodeId      any //
	InIp          any //
	OutNodeId     any //
	OutIp         any //
	Type          any //
	Protocol      any //
	Flow          any //
	TcpListenAddr any //
	UdpListenAddr any //
	InterfaceName any //
	CreatedTime   any //
	UpdatedTime   any //
	Status        any //
}
