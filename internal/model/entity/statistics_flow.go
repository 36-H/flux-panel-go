// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// StatisticsFlow is the golang structure for table statistics_flow.
type StatisticsFlow struct {
	Id          int    `json:"id"          orm:"id"           description:""` //
	UserId      int    `json:"userId"      orm:"user_id"      description:""` //
	Flow        int64  `json:"flow"        orm:"flow"         description:""` //
	TotalFlow   int64  `json:"totalFlow"   orm:"total_flow"   description:""` //
	Time        string `json:"time"        orm:"time"         description:""` //
	CreatedTime int64  `json:"createdTime" orm:"created_time" description:""` //
}
