// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2TransactionRecord is the golang structure for table v2_transaction_record.
type V2TransactionRecord struct {
	Id        int         `json:"id"         orm:"id"         ` //
	Amount    float64     `json:"amount"     orm:"amount"     ` // 交易金额
	UserId    int         `json:"user_id"    orm:"user_id"    ` // 用户ID
	Status    int         `json:"status"     orm:"status"     ` // 状态，1：交易中，2：交易完成，0：交易失败
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` // 创建时间
	Code      int         `json:"code"       orm:"code"       ` // 验证码
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` // 更新时间
}
