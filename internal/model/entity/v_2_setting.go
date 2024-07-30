// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Setting is the golang structure for table v2_setting.
type V2Setting struct {
	Code      string      `json:"code"       orm:"code"       ` //
	Value     string      `json:"value"      orm:"value"      ` //
	OrderId   int         `json:"order_id"   orm:"order_id"   ` // 顺序
	Remarks   string      `json:"remarks"    orm:"remarks"    ` // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` //
}
