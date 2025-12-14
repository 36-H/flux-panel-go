// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	Id            int    `json:"id"            orm:"id"              description:""` //
	User          string `json:"user"          orm:"user"            description:""` //
	Pwd           string `json:"pwd"           orm:"pwd"             description:""` //
	RoleId        int    `json:"roleId"        orm:"role_id"         description:""` //
	ExpTime       int64  `json:"expTime"       orm:"exp_time"        description:""` //
	Flow          int64  `json:"flow"          orm:"flow"            description:""` //
	InFlow        int64  `json:"inFlow"        orm:"in_flow"         description:""` //
	OutFlow       int64  `json:"outFlow"       orm:"out_flow"        description:""` //
	FlowResetTime int64  `json:"flowResetTime" orm:"flow_reset_time" description:""` //
	Num           int    `json:"num"           orm:"num"             description:""` //
	CreatedTime   int64  `json:"createdTime"   orm:"created_time"    description:""` //
	UpdatedTime   int64  `json:"updatedTime"   orm:"updated_time"    description:""` //
	Status        int    `json:"status"        orm:"status"          description:""` //
}
