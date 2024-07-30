// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// V2TicketMessage is the golang structure for table v2_ticket_message.
type V2TicketMessage struct {
	Id        int         `json:"id"         orm:"id"         ` //
	UserId    int         `json:"user_id"    orm:"user_id"    ` //
	TicketId  int         `json:"ticket_id"  orm:"ticket_id"  ` //
	Message   string      `json:"message"    orm:"message"    ` //
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" ` //
}
