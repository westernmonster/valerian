package cloudauth

import (
	"encoding/json"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
)

const (
	EndPoint           = "cloudauth.aliyuncs.com"
	GetVerifyToken     = "GetVerifyToken"
	GetStatus          = "GetStatus"
	SubmitVerification = "SubmitVerification"
	GetMaterials       = "GetMaterials"
	Version            = "2018-09-16"
)

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

type GetVerifyTokenResponse struct {
	Code    string          `json:"Code"`
	Success bool            `json:"Success"`
	Data    VerifyTokenData `json:"Data"`
}

type StatusData struct {
	// 认证任务所处的认证状态，取值：
	// -1： 未认证。表示没有提交记录。
	// 0： 认证中。表示已提交认证，系统正在审核认证资料。
	// 1： 认证通过。表示最近一次提交的认证资料已通过审核，当前认证任务完结。
	// 2： 认证不通过。表示最近一次提交的认证资料未通过审核，当前认证任务还可以继续发起提交。
	StatusCode int `json:"StatusCode"`
	// 认证过程中所提交的人脸照片和权威数据的比对分，取值范围为[0,100]。
	// 置信度阈值请参考：
	// 误识率0.001%时，对应阈值95。
	// 误识率0.01%时，对应阈值90。
	// 误识率为0.1%时，对应阈值80。
	// 误识率为1%时，对应阈值为60。
	// --------------------------
	// 该字段只表示人脸与权威数据的比对结果，是个参考分，通常不建议业务上仅以该分数作为是否通过的标准。认证的综合结果请参考
	// StatusCode字段， StatusCode的结果综合了人脸与权威数据的比对和其他多种策略，可以提高安全水位。
	AuthorityComparisonScore float64 `json:"AuthorityComparisonScore"`
	// 认证过程中所提交的人脸照片和身份证上的头像的相似程度分值。取值范围为[0,100]，分数越大相似度越高。
	// -------------------------
	// 只有提交的认证资料中同时包含人脸照片和身份证人像面照片，该值才有意义。
	SimilarityScore float64 `json:"SimilarityScore"`

	// 认证状态为“认证不通过”时的原因描述。不通过原因包括但不限于以下几种情况：
	// 身份证照片模糊，光线问题造成字体无法识别。
	// 身份证照片信息与认证提交信息不一致。
	// 提交的照片非身份证照片。建议您请按引导提交本人有效身份证照片。
	AuditConclusions string `json:"AuditConclusions"`
}

type GetStatusResponse struct {
	Code    string     `json:"Code"`
	Success bool       `json:"Success"`
	Data    StatusData `json:"Data"`
}

type SubmitVerificationResponse struct {
	Code    string     `json:"Code"`
	Success bool       `json:"Success"`
	Data    StatusData `json:"Data"`
}

type MaterialData struct {
	// 姓名
	Name string `json:"Name"`
	// 证件号
	IdentificationNumber string `json:"IdentificationNumber"`
	// 证件类型。identityCard代表身份证。
	IdCardType string `json:"IdCardType"`
	// 证件有效期起始日期。格式：yyyyMMdd。依赖于算法识别，非必返回字段，可能存在无法识别的情况。
	IdCardStartDate string `json:"IdCardStartDate"`
	// 证件有效期截止日期。格式：yyyyMMdd。依赖于算法识别，非必返回字段，可能存在无法识别的情况。
	IdCardExpiry string `json:"IdCardExpiry"`
	// 证件地址。通过OCR算法识别出的地址（JSON 数据格式），结构描述见Address 参数示例。
	Address string `json:"Address"`
	// 证件性别。返回值的集合为{m, f}，其中m代表男性，f代表女性。
	Sex string `json:"Sex"`
	// 证件照正面图片HTTP地址（调用本接口后1小时内可访问）。若证件类型为身份证，则为身份证人像面照片。
	IdCardFrontPic string `json:"IdCardFrontPic"`
	// 证件照背面图片HTTP地址（调用本接口后1小时内可访问）。若证件类型为身份证，则为身份证国徽面照片。
	IdCardBackPic string `json:"IdCardBackPic"`
	// 认证过程中拍摄的人像正面照图片HTTP地址（调用本接口后1小时内可访问）。
	FacePic string `json:"FacePic"`
	// 证件上的民族。依赖于算法识别，非必返回字段，可能存在无法识别的情况。
	EthnicGroup string `json:"EthnicGroup "`
}

type GetMaterialsResponse struct {
	Code    string       `json:"Code"`
	Success bool         `json:"Success"`
	Data    MaterialData `json:"Data"`
}

type CloudAuthClient struct {
	Client *sdk.Client
}

type Material struct {
	// 取值
	// Name, 姓名
	// IdentificationNumber, 身份证号
	// IdCardFrontPic, 身份证人像面照片 oss://<STSToken.bucketName>:<path_to_file>
	// IdCardBackPic, 身份证国徽面照片 oss://<STSToken.bucketName>:<path_to_file>
	// Mobile, 手机号
	Type string `json:"Type"`
}

func (p *CloudAuthClient) GetVerifyToken(ticketID string) (resp *GetVerifyTokenResponse, err error) {
	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = GetVerifyToken
	request.QueryParams["TicketId"] = ticketID
	// 使用实人认证服务的业务场景
	request.QueryParams["Biz"] = "flywiki"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["Action"] = GetVerifyToken

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":    "cloudauth",
			"method":    "GetVerifyToken",
			"biz":       "flywiki",
			"ticket_id": ticketID,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("发起实人认证请求失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetVerifyToken",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("发起实人认证请求失败")
		return
	}
	data := response.GetHttpContentBytes()
	resp = new(GetVerifyTokenResponse)
	err = json.Unmarshal(data, resp)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetVerifyToken",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("发起实人认证请求失败")
		return
	}

	if !resp.Success {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetVerifyToken",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Code: %v", resp.Code))
		err = tracerr.Errorf("发起实人认证请求失败")
		return
	}

	return
}

func (p *CloudAuthClient) GetStatus(ticketID string) (resp *GetStatusResponse, err error) {
	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = GetStatus
	request.QueryParams["TicketId"] = ticketID
	// 使用实人认证服务的业务场景
	request.QueryParams["Biz"] = "flywiki"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["Action"] = GetStatus

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":    "cloudauth",
			"method":    "GetStatus",
			"biz":       "flywiki",
			"ticket_id": ticketID,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("获取实人认证状态失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetStatus",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("获取实人认证状态失败")
		return
	}
	data := response.GetHttpContentBytes()
	resp = new(GetStatusResponse)
	err = json.Unmarshal(data, resp)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetStatus",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("获取实人认证状态失败")
		return
	}

	if !resp.Success {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetVerifyToken",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Code: %v", resp.Code))
		err = tracerr.Errorf("获取实人认证状态失败")
		return
	}

	return
}

func (p *CloudAuthClient) SubmitVerification(ticketID string, realName, idcardNumber, idcardFrontImage,
	idcardBackImage string) (resp *SubmitVerificationResponse, err error) {

	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = SubmitVerification
	request.Method = "POST"
	request.QueryParams["Action"] = SubmitVerification
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.FormParams["TicketId"] = ticketID
	// 使用实人认证服务的业务场景
	request.FormParams["Biz"] = "flywiki"
	request.FormParams["Material.1.MaterialType"] = "Name"
	request.FormParams["Material.2.MaterialType"] = "IdentificationNumber"
	request.FormParams["Material.3.MaterialType"] = "IdCardFrontPic"
	request.FormParams["Material.4.MaterialType"] = "IdCardBackPic"
	request.FormParams["Material.1.Value"] = realName
	request.FormParams["Material.2.Value"] = idcardNumber
	request.FormParams["Material.3.Value"] = idcardFrontImage
	request.FormParams["Material.4.Value"] = idcardBackImage

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":    "cloudauth",
			"method":    "SubmitVerification",
			"biz":       "flywiki",
			"ticket_id": ticketID,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("实人认证请求失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "SubmitVerification",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("实人认证请求失败")
		return
	}
	data := response.GetHttpContentBytes()
	resp = new(SubmitVerificationResponse)
	err = json.Unmarshal(data, resp)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "SubmitVerification",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("实人认证请求失败")
		return
	}

	if !resp.Success {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetVerifyToken",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Code: %v", resp.Code))
		err = tracerr.Errorf("实人认证请求失败")
		return
	}

	return
}

func (p *CloudAuthClient) GetMaterials(ticketID string) (resp *GetMaterialsResponse, err error) {
	request := requests.NewCommonRequest()
	request.Domain = EndPoint
	request.Version = Version
	request.ApiName = GetMaterials
	request.QueryParams["TicketId"] = ticketID
	// 使用实人认证服务的业务场景
	request.QueryParams["Biz"] = "flywiki"
	request.QueryParams["RegionId"] = "cn-hangzhou"
	request.QueryParams["Action"] = GetMaterials

	response, err := p.Client.ProcessCommonRequest(request)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":    "cloudauth",
			"method":    "GetMaterials",
			"biz":       "flywiki",
			"ticket_id": ticketID,
		}).Error(fmt.Sprintf("ProcessCommonRequest: %v", err))
		err = tracerr.Errorf("获取实人认证资料失败")
		return
	}

	if !response.IsSuccess() {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetMaterials",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("HTTP Status: %v", err))

		err = tracerr.Errorf("获取实人认证资料失败")
		return
	}
	data := response.GetHttpContentBytes()
	resp = new(GetMaterialsResponse)
	err = json.Unmarshal(data, resp)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetMaterials",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Unmarshal Response: %v", err))
		err = tracerr.Errorf("获取实人认证资料失败")
		return
	}

	if !resp.Success {
		logrus.WithFields(logrus.Fields{
			"prefix":       "cloudauth",
			"method":       "GetVerifyToken",
			"biz":          "flywiki",
			"ticket_id":    ticketID,
			"http_status":  response.GetHttpStatus(),
			"http_content": response.GetHttpContentString(),
		}).Error(fmt.Sprintf("Code: %v", resp.Code))
		err = tracerr.Errorf("获取实人认证资料失败")
		return
	}

	return
}
