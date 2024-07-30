// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// V2TransactionRecord is the golang structure of table v2_transaction_record for DAO operations like Where/Data.
type V2TransactionRecord struct {
	g.Meta    `orm:"table:v2_transaction_record, do:true"`
	Id        interface{} //
	Amount    interface{} // 交易金额
	UserId    interface{} // 用户ID
	Status    interface{} // 状态，1：交易中，2：交易完成，0：交易失败
	CreatedAt *gtime.Time // 创建时间
	Code      interface{} // 验证码
	UpdatedAt *gtime.Time // 更新时间
}
