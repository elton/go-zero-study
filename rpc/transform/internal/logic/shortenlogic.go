package logic

import (
	"context"

	"github.com/elton/go-zero-study/rpc/transform/internal/svc"
	"github.com/elton/go-zero-study/rpc/transform/model"

	transform "github.com/elton/go-zero-study/rpc/transform/pb"

	"github.com/tal-tech/go-zero/core/hash"
	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *transform.ShortenReq) (*transform.ShortenResp, error) {
	key := hash.Md5Hex([]byte(in.Url))[:6]

	_, err := l.svcCtx.Model.Insert(model.Shorturl{
		Shorten: key,
		Url:     in.Url,
	})

	if err != nil {
		return nil, err
	}

	return &transform.ShortenResp{
		Shorten: key,
	}, nil
}
