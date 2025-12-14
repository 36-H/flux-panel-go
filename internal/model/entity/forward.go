// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Forward is the golang structure for table forward.
type Forward struct {
	Id            int    `json:"id"            orm:"id"             description:""` //
	UserId        int    `json:"userId"        orm:"user_id"        description:""` //
	UserName      string `json:"userName"      orm:"user_name"      description:""` //
	Name          string `json:"name"          orm:"name"           description:""` //
	TunnelId      int    `json:"tunnelId"      orm:"tunnel_id"      description:""` //
	InPort        int    `json:"inPort"        orm:"in_port"        description:""` //
	OutPort       int    `json:"outPort"       orm:"out_port"       description:""` //
	RemoteAddr    string `json:"remoteAddr"    orm:"remote_addr"    description:""` //
	Strategy      string `json:"strategy"      orm:"strategy"       description:""` //
	InterfaceName string `json:"interfaceName" orm:"interface_name" description:""` //
	InFlow        int64  `json:"inFlow"        orm:"in_flow"        description:""` //
	OutFlow       int64  `json:"outFlow"       orm:"out_flow"       description:""` //
	CreatedTime   int64  `json:"createdTime"   orm:"created_time"   description:""` //
	UpdatedTime   int64  `json:"updatedTime"   orm:"updated_time"   description:""` //
	Status        int    `json:"status"        orm:"status"         description:""` //
	Inx           int    `json:"inx"           orm:"inx"            description:""` //
}
