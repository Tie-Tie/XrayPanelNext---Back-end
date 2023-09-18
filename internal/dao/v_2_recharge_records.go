// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2RechargeRecordsDao is internal type for wrapping internal DAO implements.
type internalV2RechargeRecordsDao = *internal.V2RechargeRecordsDao

// v2RechargeRecordsDao is the data access object for table v2_recharge_records.
// You can define custom methods on it to extend its functionality as you wish.
type v2RechargeRecordsDao struct {
	internalV2RechargeRecordsDao
}

var (
	// V2RechargeRecords is globally public accessible object for table v2_recharge_records operations.
	V2RechargeRecords = v2RechargeRecordsDao{
		internal.NewV2RechargeRecordsDao(),
	}
)

// Fill with you ideas below.
