package infrastructure

import (
	b64 "encoding/base64"
	"net/http"

	"valerian/infrastructure/berr"
	"valerian/infrastructure/biz"
	"valerian/library/net/http/mars"

	log "github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

// 通用返回结果
type RespCommon struct {
	// Code 状态码
	Code int `json:"code"`
	// 是否成功
	Success bool `json:"success"`
	// 返回消息
	Message string `json:"message"`
	// 返回内容，这里是一个接口字段，具体需要查看执行结果
	Result interface{} `json:"result,omitempty"`
}

type Pamarsation struct {
	// 总计数量
	Total int `json:"total"`
	// 每页数量
	PageSize int `json:"pageSize"`
	// 当前页数
	Current int `json:"current"`
}

type PamarsationResult struct {
	// 数据列表
	List interface{} `json:"list,omitempty"`
	// 分页信息
	Pamarsation Pamarsation `json:"pamarsation"`
}

// swagger:model
// 分页返回结果
type RespPamarsation struct {
	// Code 状态码
	Code int `json:"code"`
	// 是否成功
	Success bool `json:"success"`
	// 返回消息
	Message string `json:"message"`
	// 分页返回结果
	Result PamarsationResult `json:"result"`
}

type BaseCtrl struct {
}

// SuccessResp JSON 通用结果返回
func (p *BaseCtrl) SuccessResp(ctx *mars.Context, result interface{}) {
	ctx.JSON(RespCommon{
		Code:    http.StatusOK,
		Success: true,
		Message: "ok",
		Result:  result,
	}, nil)
}

// PamarsationResp JSON 分页结果
func (p *BaseCtrl) PamarsationResp(ctx *mars.Context, list interface{}, total, page, pageSize int) {
	ctx.JSON(RespPamarsation{
		Code:    http.StatusOK,
		Success: true,
		Result: PamarsationResult{
			List: list,
			Pamarsation: Pamarsation{
				Current:  page,
				Total:    total,
				PageSize: pageSize,
			},
		},
	}, nil)
}

// HandleError 错误处理
func (p *BaseCtrl) HandleError(ctx *mars.Context, err error) {
	// mode := viper.Get("MODE")
	// message := ""
	switch v := err.(type) {
	case tracerr.Error:
		// message = v.Error()
		// if mode == "development" {
		// 	tracerr.PrintSourceColor(v, 5)
		// }
		log.Error(tracerr.Sprint(v))
		break
	case *berr.BizError:
		ctx.JSON(nil, err)
		return
	default:
		// message = err.Error()
		log.Error(v)
	}
	ctx.JSON(nil, err)
	return
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
