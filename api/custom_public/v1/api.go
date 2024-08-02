package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gov2panel/internal/model/entity"
)

type PlanReq struct {
	g.Meta `path:"/plan" tags:"Captcha" method:"get" summary:"获取套餐列表"`
}
type PlanRes struct {
	g.Meta `mime:"application/json" example:"string"`
	Data   []*entity.V2Plan `json:"data"`
}

type SettingReq struct {
	g.Meta     `path:"/setting" tags:"Wallet" method:"get" summary:"充值"`
	ConfigName string `json:"config_name"`
}
type SettingRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Config string `json:"config"`
}
