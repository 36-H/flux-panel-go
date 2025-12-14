// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserTunnel is the golang structure for table user_tunnel.
type UserTunnel struct {
	Id            int   `json:"id"            orm:"id"              description:""` //
	UserId        int   `json:"userId"        orm:"user_id"         description:""` //
	TunnelId      int   `json:"tunnelId"      orm:"tunnel_id"       description:""` //
	SpeedId       int   `json:"speedId"       orm:"speed_id"        description:""` //
	Num           int   `json:"num"           orm:"num"             description:""` //
	Flow          int64 `json:"flow"          orm:"flow"            description:""` //
	InFlow        int64 `json:"inFlow"        orm:"in_flow"         description:""` //
	OutFlow       int64 `json:"outFlow"       orm:"out_flow"        description:""` //
	FlowResetTime int64 `json:"flowResetTime" orm:"flow_reset_time" description:""` //
	ExpTime       int64 `json:"expTime"       orm:"exp_time"        description:""` //
	Status        int   `json:"status"        orm:"status"          description:""` //
}
