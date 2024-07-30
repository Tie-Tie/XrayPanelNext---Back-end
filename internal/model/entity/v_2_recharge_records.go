// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2RechargeRecords is the golang structure for table v2_recharge_records.
type V2RechargeRecords struct {
	Id              int         `json:"id"               orm:"id"               ` //
	Amount          float64     `json:"amount"           orm:"amount"           ` // 金额
	UserId          int         `json:"user_id"          orm:"user_id"          ` // 用户id
	OperateType     int         `json:"operate_type"     orm:"operate_type"     ` // 1充值 2消费
	RechargeName    string      `json:"recharge_name"    orm:"recharge_name"    ` // 充值类型 operate_type=1才有
	ConsumptionName string      `json:"consumption_name" orm:"consumption_name" ` // 消费类型 operate_type=2才有
	Remarks         string      `json:"remarks"          orm:"remarks"          ` // 备注
	TransactionId   string      `json:"transaction_id"   orm:"transaction_id"   ` // 订单号 规则看程序注释
	CreatedAt       *gtime.Time `json:"created_at"       orm:"created_at"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updated_at"       orm:"updated_at"       ` // 更新时间
	Status          int         `json:"status"           orm:"status"           ` // 订单状态：0：失败，1：等待中，2：成功
	Code            int         `json:"code"             orm:"code"             ` // 验证码
	RechargeMethod  int         `json:"recharge_method"  orm:"recharge_method"  ` // 充值方式：0：ERC20 USDT 、1：TRC20 USDT
}
