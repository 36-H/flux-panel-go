// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForwardDao is the data access object for the table forward.
type ForwardDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ForwardColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ForwardColumns defines and stores column names for the table forward.
type ForwardColumns struct {
	Id            string //
	UserId        string //
	UserName      string //
	Name          string //
	TunnelId      string //
	InPort        string //
	OutPort       string //
	RemoteAddr    string //
	Strategy      string //
	InterfaceName string //
	InFlow        string //
	OutFlow       string //
	CreatedTime   string //
	UpdatedTime   string //
	Status        string //
	Inx           string //
}

// forwardColumns holds the columns for the table forward.
var forwardColumns = ForwardColumns{
	Id:            "id",
	UserId:        "user_id",
	UserName:      "user_name",
	Name:          "name",
	TunnelId:      "tunnel_id",
	InPort:        "in_port",
	OutPort:       "out_port",
	RemoteAddr:    "remote_addr",
	Strategy:      "strategy",
	InterfaceName: "interface_name",
	InFlow:        "in_flow",
	OutFlow:       "out_flow",
	CreatedTime:   "created_time",
	UpdatedTime:   "updated_time",
	Status:        "status",
	Inx:           "inx",
}

// NewForwardDao creates and returns a new DAO object for table data access.
func NewForwardDao(handlers ...gdb.ModelHandler) *ForwardDao {
	return &ForwardDao{
		group:    "default",
		table:    "forward",
		columns:  forwardColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForwardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForwardDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForwardDao) Columns() ForwardColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForwardDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForwardDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForwardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
