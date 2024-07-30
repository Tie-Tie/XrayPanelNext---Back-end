// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2InvitationRecords is the golang structure for table v2_invitation_records.
type V2InvitationRecords struct {
	Id                int         `json:"id"                  orm:"id"                  ` //
	Amount            float64     `json:"amount"              orm:"amount"              ` // 金额
	UserId            int         `json:"user_id"             orm:"user_id"             ` // 邀请者
	FromUserId        int         `json:"from_user_id"        orm:"from_user_id"        ` // 被邀请者
	CommissionRate    int         `json:"commission_rate"     orm:"commission_rate"     ` // 佣金比例
	RechargeRecordsId int         `json:"recharge_records_id" orm:"recharge_records_id" ` // 订单id
	CreatedAt         *gtime.Time `json:"created_at"          orm:"created_at"          ` // 创建时间
	UpdatedAt         *gtime.Time `json:"updated_at"          orm:"updated_at"          ` // 更新时间
	OperateType       int         `json:"operate_type"        orm:"operate_type"        ` // 1邀请 2提现
	State             int         `json:"state"               orm:"state"               ` // 状态 0未审核 1审核 2拒绝
}
