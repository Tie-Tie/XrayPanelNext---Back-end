package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/model/model"
)

type BuyReq struct {
	g.Meta  `path:"/buy" tags:"Custom_User" method:"post" summary:"购买"`
	PlanId  int    `json:"plan_id"` //订阅id
	Code    string `json:"code"`    //优惠码
	TUserID int    //用户id
}
type BuyRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type PlanInfo struct {
	Token          string            `json:"token"`               //订阅token
	TransferEnable int64             `json:"transfer_enable"    ` //总流量
	Data           *entity.V2Plan    `json:"data"    `            //总流量
	U              int64             `json:"u"`
	D              int64             `json:"d"`
	ExpiredAt      *gtime.Time       `json:"expired_at"         ` //到期时间
	PlanName       string            `json:"plan_name"`           //订阅名
	UserName       string            `json:"user_name"`           //用户名
	Setting        map[string]*g.Var `json:"setting"`             //设置
}
type IndexReq struct {
	g.Meta `path:"/index" tags:"Custom_User" method:"get" summary:"获取用户首页数据"`
}
type IndexRes struct {
	g.Meta          `mime:"text/html" example:"string"`
	SubscribeDomain string    `json:"subscribe_domain"`
	UserPageHtml    string    `json:"user_page_html"`
	PlanInfo        *PlanInfo `json:"plan_info"`
}

type KnowledgeReq struct {
	g.Meta `path:"/knowledge" tags:"Knowledge" method:"get" summary:"使用文档"`
	entity.V2Knowledge
}
type KnowledgeRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Data   []*model.KnowledgeInfo `json:"data"`
}

type WalletReq struct {
	g.Meta  `path:"/wallet" tags:"Wallet" method:"get,post" summary:"钱包页面"`
	TUserID int
}
type WalletRes struct {
	g.Meta      `mime:"text/html" example:"string"`
	User        entity.V2User `json:"user"`
	InviteCount int           `json:"invite_count"`
	CType       int           `json:"ctype"`
	CRate       int           `json:"crate"`
}

type TopUpReq struct {
	g.Meta         `path:"/topUp" tags:"Wallet" method:"post" summary:"充值"`
	Amount         float64 `json:"amount"`
	RechargeMethod int     `json:"recharge_method"`
}
type TopUpRes struct {
	g.Meta     `mime:"text/html" example:"string"`
	Amount     float64 `json:"code"`
	ExpiryTime int64   `json:"expiry_time"`
	Success    bool    `json:"success"`
}
