// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2TransactionRecordDao is internal type for wrapping internal DAO implements.
type internalV2TransactionRecordDao = *internal.V2TransactionRecordDao

// v2TransactionRecordDao is the data access object for table v2_transaction_record.
// You can define custom methods on it to extend its functionality as you wish.
type v2TransactionRecordDao struct {
	internalV2TransactionRecordDao
}

var (
	// V2TransactionRecord is globally public accessible object for table v2_transaction_record operations.
	V2TransactionRecord = v2TransactionRecordDao{
		internal.NewV2TransactionRecordDao(),
	}
)

// Fill with you ideas below.