package custom_user

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"gov2panel/internal/model/entity"
	"gov2panel/internal/service"

	"gov2panel/api/custom_user/v1"
)

func (c *ControllerV1) Buy(ctx context.Context, req *v1.BuyReq) (res *v1.BuyRes, err error) {
	res = &v1.BuyRes{}

	var user entity.V2User
	err = g.RequestFromCtx(ctx).GetCtxVar("database_user").Struct(&user)
	if err != nil {
		g.RequestFromCtx(ctx).Response.Write(err.Error())
		return
	}

	//检查要购买的套餐
	plan, err := service.Plan().GetPlanById(req.PlanId)
	if err != nil {
		return
	}
	if plan == nil {
		return res, errors.New("套餐不存在")
	}
	if plan.Show != 1 {
		return res, errors.New("套餐未开启")
	}
	if plan.Price < 0 || plan.Expired < 0 {
		return res, errors.New("套餐设置不对请联系管理员")
	}

	err = service.Plan().UserBuyAndRenew(req.Code, plan, &user)
	return
}

func (c *ControllerV1) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	g.Dump(666)
	var user entity.V2User
	var setting map[string]*g.Var
	var plan *entity.V2Plan

	if err = g.RequestFromCtx(ctx).GetCtxVar("database_user").Struct(&user); err != nil {
		g.RequestFromCtx(ctx).Response.Write(err.Error())
		return
	}

	if plan, err = service.Plan().GetPlanById(user.GroupId); err != nil {
		g.RequestFromCtx(ctx).Response.Write(err.Error())
		return
	}

	if setting, err = service.Setting().GetSettingAllMap(); err != nil {
		return res, err
	}

	res = &v1.IndexRes{
		UserPageHtml:    setting["user_page_html"].String(),
		SubscribeDomain: setting["subscribe_domain"].String(),
		PlanInfo: &v1.PlanInfo{
			Token:          user.Token,
			U:              user.U,
			D:              user.D,
			ExpiredAt:      user.ExpiredAt,
			PlanName:       plan.Name,
			UserName:       user.UserName,
			TransferEnable: user.TransferEnable,
			Data:           plan,
		},
	}

	return
}

func (c *ControllerV1) Knowledge(ctx context.Context, req *v1.KnowledgeReq) (res *v1.KnowledgeRes, err error) {
	res = &v1.KnowledgeRes{}
	res.Data, err = service.Knowledge().GetKnowledgeShowList(req.V2Knowledge)
	return
}
func (c *ControllerV1) Wallet(ctx context.Context, req *v1.WalletReq) (res *v1.WalletRes, err error) {
	res = &v1.WalletRes{}

	var inviteCount int
	if inviteCount, err = service.User().GetInviteCountByUserId(req.TUserID); err != nil {
		return res, err
	}

	var user entity.V2User
	if err = g.RequestFromCtx(ctx).GetCtxVar("database_user").Struct(&user); err != nil {
		return nil, err
	}

	cType, cRate := service.User().GetUserCTypeAndCRate(&user)

	res.CType = cType
	res.CRate = cRate
	res.User = user
	res.InviteCount = inviteCount

	return
}
