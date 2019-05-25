package infrastructure

import (
	b64 "encoding/base64"

	"valerian/infrastructure/biz"
	"valerian/library/net/http/mars"

	"github.com/ztrue/tracerr"
)

type BaseCtrl struct {
}

// Base64Decode Base64 编码
func (p *BaseCtrl) Base64Decode(b64Str string) (result string, err error) {
	l, err := b64.URLEncoding.DecodeString(b64Str)
	return string(l), err
}

// GetAccountID 获取用户ID
func (p *BaseCtrl) GetAccountID(ctx *mars.Context) (accountID int64, err error) {
	aid, exist := ctx.Get("AccountID")
	if !exist {
		err = tracerr.New("获取当前用户信息失败")
		return
	}

	accountID = aid.(int64)
	return
}

// GetBizContext 获取Biz上下文
func (p *BaseCtrl) GetBizContext(ctx *mars.Context) (bizContext *biz.BizContext) {
	bizContext = &biz.BizContext{
		Locale: "zh-CN",
	}

	accountID, _ := p.GetAccountID(ctx)
	if accountID != 0 {
		bizContext.AccountID = &accountID
	}

	locale := ctx.Request.Header.Get("Locale")

	for _, v := range []string{"zh-CN", "en-US"} {
		if locale == v {
			bizContext.Locale = v
			break
		}
	}

	return
}
