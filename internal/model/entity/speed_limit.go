// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SpeedLimit is the golang structure for table speed_limit.
type SpeedLimit struct {
	Id          int    `json:"id"          orm:"id"           description:""` //
	Name        string `json:"name"        orm:"name"         description:""` //
	Speed       int    `json:"speed"       orm:"speed"        description:""` //
	TunnelId    int    `json:"tunnelId"    orm:"tunnel_id"    description:""` //
	TunnelName  string `json:"tunnelName"  orm:"tunnel_name"  description:""` //
	CreatedTime int64  `json:"createdTime" orm:"created_time" description:""` //
	UpdatedTime int64  `json:"updatedTime" orm:"updated_time" description:""` //
	Status      int    `json:"status"      orm:"status"       description:""` //
}
