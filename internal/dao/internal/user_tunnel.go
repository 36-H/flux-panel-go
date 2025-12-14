// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserTunnelDao is the data access object for the table user_tunnel.
type UserTunnelDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserTunnelColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserTunnelColumns defines and stores column names for the table user_tunnel.
type UserTunnelColumns struct {
	Id            string //
	UserId        string //
	TunnelId      string //
	SpeedId       string //
	Num           string //
	Flow          string //
	InFlow        string //
	OutFlow       string //
	FlowResetTime string //
	ExpTime       string //
	Status        string //
}

// userTunnelColumns holds the columns for the table user_tunnel.
var userTunnelColumns = UserTunnelColumns{
	Id:            "id",
	UserId:        "user_id",
	TunnelId:      "tunnel_id",
	SpeedId:       "speed_id",
	Num:           "num",
	Flow:          "flow",
	InFlow:        "in_flow",
	OutFlow:       "out_flow",
	FlowResetTime: "flow_reset_time",
	ExpTime:       "exp_time",
	Status:        "status",
}

// NewUserTunnelDao creates and returns a new DAO object for table data access.
func NewUserTunnelDao(handlers ...gdb.ModelHandler) *UserTunnelDao {
	return &UserTunnelDao{
		group:    "default",
		table:    "user_tunnel",
		columns:  userTunnelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserTunnelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserTunnelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserTunnelDao) Columns() UserTunnelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserTunnelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserTunnelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserTunnelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
