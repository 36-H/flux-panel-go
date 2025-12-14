// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Node is the golang structure of table node for DAO operations like Where/Data.
type Node struct {
	g.Meta      `orm:"table:node, do:true"`
	Id          any //
	Name        any //
	Secret      any //
	Ip          any //
	ServerIp    any //
	PortSta     any //
	PortEnd     any //
	Version     any //
	Http        any //
	Tls         any //
	Socks       any //
	CreatedTime any //
	UpdatedTime any //
	Status      any //
}
