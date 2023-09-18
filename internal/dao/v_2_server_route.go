// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2ServerRouteDao is internal type for wrapping internal DAO implements.
type internalV2ServerRouteDao = *internal.V2ServerRouteDao

// v2ServerRouteDao is the data access object for table v2_server_route.
// You can define custom methods on it to extend its functionality as you wish.
type v2ServerRouteDao struct {
	internalV2ServerRouteDao
}

var (
	// V2ServerRoute is globally public accessible object for table v2_server_route operations.
	V2ServerRoute = v2ServerRouteDao{
		internal.NewV2ServerRouteDao(),
	}
)

// Fill with you ideas below.
