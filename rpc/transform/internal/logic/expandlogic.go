package logic

import (
	"context"

	transform "github.com/elton/go-zero-study/rpc/transform/pb"

	"github.com/elton/go-zero-study/rpc/transform/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *transform.ExpandReq) (*transform.ExpandResp, error) {
	res, err := l.svcCtx.Model.FindOne(in.Shorten)
	if err != nil {
		return nil, err
	}

	return &transform.ExpandResp{
		Url: res.Url,
	}, nil
}
