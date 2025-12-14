// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserTunnel is the golang structure of table user_tunnel for DAO operations like Where/Data.
type UserTunnel struct {
	g.Meta        `orm:"table:user_tunnel, do:true"`
	Id            any //
	UserId        any //
	TunnelId      any //
	SpeedId       any //
	Num           any //
	Flow          any //
	InFlow        any //
	OutFlow       any //
	FlowResetTime any //
	ExpTime       any //
	Status        any //
}
