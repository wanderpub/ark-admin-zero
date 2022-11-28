package handlers

import (
	"ark-admin-zero/app/core/model"
	"context"
)

type IHandler interface {
	DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error)
}
