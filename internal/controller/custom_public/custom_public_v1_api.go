package custom_public

import (
	"context"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"

	"gov2panel/api/custom_public/v1"
)

func (c *ControllerV1) Plan(ctx context.Context, req *v1.PlanReq) (res *v1.PlanRes, err error) {
	var querySearch entity.V2Plan
	list, err := service.Plan().GetPlanAllList(querySearch)
	if err != nil {
		return nil, err
	}
	res = &v1.PlanRes{
		Data: list,
	}
	return
}
