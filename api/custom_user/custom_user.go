// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package custom_user

import (
	"context"

	"gov2panel/api/custom_user/v1"
)

type ICustomUserV1 interface {
	Buy(ctx context.Context, req *v1.BuyReq) (res *v1.BuyRes, err error)
	Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error)
	Knowledge(ctx context.Context, req *v1.KnowledgeReq) (res *v1.KnowledgeRes, err error)
	Wallet(ctx context.Context, req *v1.WalletReq) (res *v1.WalletRes, err error)
}
