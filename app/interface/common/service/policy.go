package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"hash"
	"io"
	"path/filepath"
	"strconv"
	"time"

	"valerian/app/interface/common/model"
	"valerian/library/gid"
)

const (
	// Host             = "https://flywiki.oss-cn-hangzhou.aliyuncs.com"
	// CallbackURL      = "https://dev.flywk.com/api/v1/files/callback"
	// ImageDir         = "images/"
	// FileDir          = "files/"
	// CertificationDir = "certifications/"
	// OtherDir         = "other/"
	ExpireTime = int64(60)

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

func (p *Service) GetPolicyToken(fileType, fileName string) (token model.PolicyToken, err error) {
	dir := p.c.OSS.OtherDir
	switch fileType {
	case "image":
		dir = p.c.OSS.ImageDir
		break
	case "certification":
		dir = p.c.OSS.CertificationDir
		break
	case "file":
		dir = p.c.OSS.FileDir
		break
	}

	id := gid.NewID()

	ext := filepath.Ext(fileName)
	name := strconv.FormatInt(id, 10)

	if ext != "" {
		name = name + ext
	}

	now := time.Now().Unix()
	expire_end := now + ExpireTime
	var tokenExpire = get_gmt_iso8601(expire_end)

	//create post policy json
	var config model.ConfigStruct
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
		return
	}

	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(p.c.Aliyun.AccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam model.CallbackParam
	callbackParam.CallbackUrl = p.c.OSS.CallbackURL
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callback_str, err := json.Marshal(callbackParam)
	if err != nil {
		return
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callback_str)

	var policyToken model.PolicyToken
	policyToken.AccessKeyId = p.c.Aliyun.AccessKeyID
	policyToken.Host = p.c.OSS.Host
	policyToken.Expire = expire_end
	policyToken.Signature = string(signedStr)
	policyToken.Directory = dir
	policyToken.Policy = string(debyte)
	policyToken.Callback = string(callbackBase64)
	policyToken.Key = dir + name

	token = policyToken
	return
}
