// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package custom_public

import (
	"context"

	"gov2panel/api/custom_public/v1"
)

type ICustomPublicV1 interface {
	Plan(ctx context.Context, req *v1.PlanReq) (res *v1.PlanRes, err error)
	Setting(ctx context.Context, req *v1.SettingReq) (res *v1.SettingRes, err error)
}
