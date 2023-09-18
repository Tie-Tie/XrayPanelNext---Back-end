// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Plan is the golang structure for table v2_plan.
type V2Plan struct {
	Id                 int         `json:"id"                   ` //
	TransferEnable     float64     `json:"transfer_enable"      ` // 流量(GB)
	SpeedLimit         int         `json:"speed_limit"          ` // 速度限制
	Name               string      `json:"name"                 ` // 名称
	Show               int         `json:"show"                 ` // 是否显示
	OrderId            int         `json:"order_id"             ` // 顺序
	Renew              int         `json:"renew"                ` // 是否允许续购
	Content            string      `json:"content"              ` // 描述
	Expired            int         `json:"expired"              ` // 有效期 day
	Price              float64     `json:"price"                ` // 价格
	ResetTrafficMethod int         `json:"reset_traffic_method" ` // 套餐类型，1 覆盖、2 叠加
	CapacityLimit      int         `json:"capacity_limit"       ` // 最大用户
	CreatedAt          *gtime.Time `json:"created_at"           ` //
	UpdatedAt          *gtime.Time `json:"updated_at"           ` //
	Remarks            string      `json:"remarks"              ` // 备注
}
