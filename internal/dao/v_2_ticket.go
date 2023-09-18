// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2TicketDao is internal type for wrapping internal DAO implements.
type internalV2TicketDao = *internal.V2TicketDao

// v2TicketDao is the data access object for table v2_ticket.
// You can define custom methods on it to extend its functionality as you wish.
type v2TicketDao struct {
	internalV2TicketDao
}

var (
	// V2Ticket is globally public accessible object for table v2_ticket operations.
	V2Ticket = v2TicketDao{
		internal.NewV2TicketDao(),
	}
)

// Fill with you ideas below.
