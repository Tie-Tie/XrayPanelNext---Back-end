// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2Knowledge is the golang structure for table v2_knowledge.
type V2Knowledge struct {
	Id        int         `json:"id"         orm:"id"         ` //
	Category  string      `json:"category"   orm:"category"   ` // 分類名
	Title     string      `json:"title"      orm:"title"      ` // 標題
	Body      string      `json:"body"       orm:"body"       ` // 內容
	OrderId   int         `json:"order_id"   orm:"order_id"   ` // 排序
	Show      int         `json:"show"       orm:"show"       ` // 顯示
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` //
}
