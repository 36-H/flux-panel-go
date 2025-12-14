// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SpeedLimitDao is the data access object for the table speed_limit.
type SpeedLimitDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SpeedLimitColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SpeedLimitColumns defines and stores column names for the table speed_limit.
type SpeedLimitColumns struct {
	Id          string //
	Name        string //
	Speed       string //
	TunnelId    string //
	TunnelName  string //
	CreatedTime string //
	UpdatedTime string //
	Status      string //
}

// speedLimitColumns holds the columns for the table speed_limit.
var speedLimitColumns = SpeedLimitColumns{
	Id:          "id",
	Name:        "name",
	Speed:       "speed",
	TunnelId:    "tunnel_id",
	TunnelName:  "tunnel_name",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
	Status:      "status",
}

// NewSpeedLimitDao creates and returns a new DAO object for table data access.
func NewSpeedLimitDao(handlers ...gdb.ModelHandler) *SpeedLimitDao {
	return &SpeedLimitDao{
		group:    "default",
		table:    "speed_limit",
		columns:  speedLimitColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SpeedLimitDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SpeedLimitDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SpeedLimitDao) Columns() SpeedLimitColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SpeedLimitDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SpeedLimitDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SpeedLimitDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
