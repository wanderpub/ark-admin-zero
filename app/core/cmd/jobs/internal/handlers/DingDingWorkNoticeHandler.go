package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"
	"fmt"
)

type dingDingWorkHandler struct {
	svcCtx *svc.ServiceContext
}

func NewDingDingWorkNoticeHandler(svcCtx *svc.ServiceContext) IHandler {
	return dingDingWorkHandler{
		svcCtx: svcCtx,
	}
}

func (h dingDingWorkHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	fmt.Println(sendTask)
	return nil
}
