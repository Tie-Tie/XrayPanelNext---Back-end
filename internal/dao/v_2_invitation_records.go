// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2InvitationRecordsDao is internal type for wrapping internal DAO implements.
type internalV2InvitationRecordsDao = *internal.V2InvitationRecordsDao

// v2InvitationRecordsDao is the data access object for table v2_invitation_records.
// You can define custom methods on it to extend its functionality as you wish.
type v2InvitationRecordsDao struct {
	internalV2InvitationRecordsDao
}

var (
	// V2InvitationRecords is globally public accessible object for table v2_invitation_records operations.
	V2InvitationRecords = v2InvitationRecordsDao{
		internal.NewV2InvitationRecordsDao(),
	}
)

// Fill with you ideas below.
