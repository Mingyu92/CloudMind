package user

import (
	"context"

	"CloudMind/app/usercenter/cmd/api/internal/svc"
	"CloudMind/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendemailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendemailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendemailLogic {
	return &SendemailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendemailLogic) Sendemail(req *types.SendEmailReq) (resp *types.SendEmailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
