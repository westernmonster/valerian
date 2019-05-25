package http

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/ztrue/tracerr"

	"valerian/infrastructure"
	"valerian/library/gid"
	"valerian/library/net/http/mars"
	"valerian/models"
)

const (
	Host             = "https://flywiki.oss-cn-hangzhou.aliyuncs.com"
	CallbackURL      = "https://dev.flywk.com/api/v1/files/callback"
	ImageDir         = "images/"
	FileDir          = "files/"
	CertificationDir = "certifications/"
	OtherDir         = "other/"
	ExpireTime       = int64(30)

	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func get_gmt_iso8601(expire_end int64) string {
	var tokenExpire = time.Unix(expire_end, 0).Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

func GetPolicyToken(fileType, fileName string) (token models.PolicyToken, err error) {
	dir := OtherDir
	switch fileType {
	case "image":
		dir = ImageDir
		break
	case "certification":
		dir = CertificationDir
		break
	case "file":
		dir = FileDir
		break
	}

	id, err := gid.NextID()
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	ext := filepath.Ext(fileName)
	name := strconv.FormatInt(id, 10)

	if ext != "" {
		name = name + ext
	}

	mode := viper.Get("MODE")
	accessKeyID := viper.GetString(fmt.Sprintf("%s.aliyun.access_key_id", mode))
	accessKeySecret := viper.GetString(fmt.Sprintf("%s.aliyun.access_key_secret", mode))

	now := time.Now().Unix()
	expire_end := now + ExpireTime
	var tokenExpire = get_gmt_iso8601(expire_end)

	//create post policy json
	var config models.ConfigStruct
	config.Expiration = tokenExpire
	config.Conditions = append(config.Conditions, []interface{}{
		"content-length-range",
		0,
		1024 * 1024 * 5,
	})

	config.Conditions = append(config.Conditions, []interface{}{
		"eq",
		"$key",
		dir + name,
	})

	//calucate signature
	result, err := json.Marshal(config)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}

	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(accessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam models.CallbackParam
	callbackParam.CallbackUrl = CallbackURL
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callback_str, err := json.Marshal(callbackParam)
	if err != nil {
		err = tracerr.Wrap(err)
		return
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callback_str)

	var policyToken models.PolicyToken
	policyToken.AccessKeyId = accessKeyID
	policyToken.Host = Host
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = dir
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)
	policyToken.Key = dir + name

	token = policyToken
	return
}

type FileCtrl struct {
	infrastructure.BaseCtrl
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
// @Param req body models.RequestOSSTokenReq true "请求"
// @Success 200 {object} models.PolicyToken "Token"
// @Failure 400 "验证失败"
// @Failure 401 "登录验证失败"
// @Failure 500 "服务器端错误"
// @Router /files/oss_token [post]
func (p *FileCtrl) GetOSSToken(ctx *mars.Context) {
	req := new(models.RequestOSSTokenReq)
	ctx.Bind(req)

	if e := req.Validate(); e != nil {
		ctx.JSON(nil, e)
	}

	token, err := GetPolicyToken(req.FileType, req.FileName)
	if err != nil {
		p.HandleError(ctx, err)
		return
	}

	p.SuccessResp(ctx, token)

	return
}
