// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TunnelDao is the data access object for the table tunnel.
type TunnelDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TunnelColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TunnelColumns defines and stores column names for the table tunnel.
type TunnelColumns struct {
	Id            string //
	Name          string //
	TrafficRatio  string //
	InNodeId      string //
	InIp          string //
	OutNodeId     string //
	OutIp         string //
	Type          string //
	Protocol      string //
	Flow          string //
	TcpListenAddr string //
	UdpListenAddr string //
	InterfaceName string //
	CreatedTime   string //
	UpdatedTime   string //
	Status        string //
}

// tunnelColumns holds the columns for the table tunnel.
var tunnelColumns = TunnelColumns{
	Id:            "id",
	Name:          "name",
	TrafficRatio:  "traffic_ratio",
	InNodeId:      "in_node_id",
	InIp:          "in_ip",
	OutNodeId:     "out_node_id",
	OutIp:         "out_ip",
	Type:          "type",
	Protocol:      "protocol",
	Flow:          "flow",
	TcpListenAddr: "tcp_listen_addr",
	UdpListenAddr: "udp_listen_addr",
	InterfaceName: "interface_name",
	CreatedTime:   "created_time",
	UpdatedTime:   "updated_time",
	Status:        "status",
}

// NewTunnelDao creates and returns a new DAO object for table data access.
func NewTunnelDao(handlers ...gdb.ModelHandler) *TunnelDao {
	return &TunnelDao{
		group:    "default",
		table:    "tunnel",
		columns:  tunnelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TunnelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TunnelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TunnelDao) Columns() TunnelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TunnelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TunnelDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TunnelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
