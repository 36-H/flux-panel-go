// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Node is the golang structure for table node.
type Node struct {
	Id          int    `json:"id"          orm:"id"           description:""` //
	Name        string `json:"name"        orm:"name"         description:""` //
	Secret      string `json:"secret"      orm:"secret"       description:""` //
	Ip          string `json:"ip"          orm:"ip"           description:""` //
	ServerIp    string `json:"serverIp"    orm:"server_ip"    description:""` //
	PortSta     int    `json:"portSta"     orm:"port_sta"     description:""` //
	PortEnd     int    `json:"portEnd"     orm:"port_end"     description:""` //
	Version     string `json:"version"     orm:"version"      description:""` //
	Http        int    `json:"http"        orm:"http"         description:""` //
	Tls         int    `json:"tls"         orm:"tls"          description:""` //
	Socks       int    `json:"socks"       orm:"socks"        description:""` //
	CreatedTime int64  `json:"createdTime" orm:"created_time" description:""` //
	UpdatedTime int64  `json:"updatedTime" orm:"updated_time" description:""` //
	Status      int    `json:"status"      orm:"status"       description:""` //
}
