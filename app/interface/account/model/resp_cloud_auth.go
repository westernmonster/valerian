package model

type VerifyToken struct {
	// 一次认证会话的标识。
	Token string `json:"Token"`
	// 认证会话超时时间，单位为秒。一般为1,800s。
	DurationSeconds int `json:"DurationSeconds"`
}

type STSToken struct {
	// OSS AccessKey ID。
	// required: true
	AccessKeyId string `json:"AccessKeyId"`
	// OSS AccessKey密钥
	AccessKeySecret string `json:"AccessKeySecret"`
	// STS token过期时间。
	Expiration string `json:"Expiration"`
	// OSS endpoint
	EndPoint string `json:"EndPoint"`
	// OSS bucket，认证服务指定的文件上传 Bucket。
	BucketName string `json:"BucketName"`
	//  一个认证请求生成的用来上传的目录，业务方需要将文件上传到这个目录。
	Path string `json:"Path"`
	// STS 访问Token。
	Token string `json:"Token"`
}
type VerifyTokenData struct {

	// 认证流程页面入口URL。
	CloudauthPageUrl string `json:"CloudauthPageUrl"`

	// 如果业务方有额外的文件需要一并提交认证服务，可以使用STSToken将其上传到认证服务指定的地方。
	// 操作方法参考使用STSToken 上传文件示例。STSToken的具体结构描述见STSToken。
	STSToken STSToken `json:"STSToken"`

	// 认证流程页面入口URL。
	VerifyToken VerifyToken `json:"VerifyToken"`
}
