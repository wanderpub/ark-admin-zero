package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"
	"fmt"
)

type enterpriseWeChatHandler struct {
	svcCtx *svc.ServiceContext
}

func NewEnterpriseWeChatHandler(svcCtx *svc.ServiceContext) IHandler {
	return enterpriseWeChatHandler{
		svcCtx: svcCtx,
	}
}

func (h enterpriseWeChatHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	fmt.Println(sendTask)
	return nil
}
