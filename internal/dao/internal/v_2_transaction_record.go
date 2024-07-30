// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// V2TransactionRecordDao is the data access object for table v2_transaction_record.
type V2TransactionRecordDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns V2TransactionRecordColumns // columns contains all the column names of Table for convenient usage.
}

// V2TransactionRecordColumns defines and stores column names for table v2_transaction_record.
type V2TransactionRecordColumns struct {
	Id        string //
	Amount    string // 交易金额
	UserId    string // 用户ID
	Status    string // 状态，1：交易中，2：交易完成，0：交易失败
	CreatedAt string // 创建时间
	Code      string // 验证码
	UpdatedAt string // 更新时间
}

// v2TransactionRecordColumns holds the columns for table v2_transaction_record.
var v2TransactionRecordColumns = V2TransactionRecordColumns{
	Id:        "id",
	Amount:    "amount",
	UserId:    "user_id",
	Status:    "status",
	CreatedAt: "created_at",
	Code:      "code",
	UpdatedAt: "updated_at",
}

// NewV2TransactionRecordDao creates and returns a new DAO object for table data access.
func NewV2TransactionRecordDao() *V2TransactionRecordDao {
	return &V2TransactionRecordDao{
		group:   "default",
		table:   "v2_transaction_record",
		columns: v2TransactionRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *V2TransactionRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *V2TransactionRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *V2TransactionRecordDao) Columns() V2TransactionRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *V2TransactionRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *V2TransactionRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *V2TransactionRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
