package http

import (
	"valerian/app/interface/common/model"
	"valerian/library/ecode"
	"valerian/library/net/http/mars"
)

// @Summary 创建Offie文档转换任务
// @Description 文档转换任务
// @Tags common
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body app.interface.common.model.ArgConvertOfficeTask true "请求"
// @Success 200 "成功"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /file/office_convert [post]
func officeConvert(c *mars.Context) {
	arg := new(model.ArgConvertOfficeTask)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}
	c.JSON(nil, srv.CreateOfficeConversionTask(c, arg))
}

// @Summary 获取阿里云OSS STS上传凭证
// @Description 获取阿里云OSS STS上传凭证
// @Tags common
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Success 200 {object}  app.interface.common.model.STSResp "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /file/sts_cred [get]
func stsCred(c *mars.Context) {
	c.JSON(srv.AssumeRole(c))
}

// @Summary 获取阿里云OSS上传TOKEN
// @Description 获取阿里云OSS上传TOKEN
// @Description 阿里云文档：https://help.aliyun.com/document_detail/31926.html
// @Tags common
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.common.model.ArgOSSToken true "请求"
// @Success 200 {object}  app.interface.common.model.PolicyToken "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /file/oss_token [post]
func ossToken(c *mars.Context) {
	arg := new(model.ArgOSSToken)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.GetPolicyToken(arg.FileType, arg.FileName))
}

// @Summary 图片URL上传到OSS
// @Description 图片URL上传到OSS
// @Tags common
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer"
// @Param Source header int true "Source 来源，1:Web, 2:iOS; 3:Android" Enums(1, 2, 3)
// @Param Locale header string true "语言" Enums(zh-CN, en-US)
// @Param req body  app.interface.common.model.ArgUploadURL true "请求"
// @Success 200 {object}  app.interface.common.model.UploadURLResp "File"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /file/url_upload [post]
func urlUpload(c *mars.Context) {
	arg := new(model.ArgUploadURL)
	if e := c.Bind(arg); e != nil {
		return
	}

	if e := arg.Validate(); e != nil {
		c.JSON(nil, ecode.RequestErr)
		return
	}

	c.JSON(srv.UploadImageURL(c, arg))
}
