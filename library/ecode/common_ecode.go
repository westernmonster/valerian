package ecode

import "github.com/pkg/errors"

// All common ecode
var (
	OK = add(200) // 正确

	AppKeyInvalid           = add(1)   // 应用程序不存在或已被封禁
	AccessKeyErr            = add(2)   // Access Key错误
	SignCheckErr            = add(3)   // API校验密匙错误
	MethodNoPermission      = add(4)   // 调用方对该Method没有权限
	NoLogin                 = add(101) // 账号未登录
	UserDisabled            = add(102) // 账号被封停
	CaptchaErr              = add(105) // 验证码错误
	UserInactive            = add(106) // 账号未激活
	AppDenied               = add(108) // 应用不存在或者被封禁
	MobileNoVerfiy          = add(110) // 未绑定手机
	CsrfNotMatchErr         = add(111) // csrf 校验失败
	ServiceUpdate           = add(112) // 系统升级中
	UserIDCheckInvalid      = add(113) // 账号尚未实名认证
	UserIDCheckInvalidPhone = add(114) // 请先绑定手机
	UserIDCheckInvalidCard  = add(115) // 请先完成实名认证
	ClientNotExist          = add(116) // Client不存在

	NotModified         = add(304) // 木有改动
	TemporaryRedirect   = add(307) // 撞车跳转
	RequestErr          = add(400) // 请求错误
	Unauthorized        = add(401) // 未认证
	AccessDenied        = add(403) // 访问权限不足
	NothingFound        = add(404) // 啥都木有
	MethodNotAllowed    = add(405) // 不支持该方法
	Conflict            = add(409) // 冲突
	ServerErr           = add(500) // 服务器错误
	ServiceUnavailable  = add(503) // 过载保护,服务暂不可用
	Deadline            = add(504) // 服务调用超时
	LimitExceed         = add(509) // 超出限制
	FileTooLarge        = add(617) // 上传文件太大
	FailedTooManyTimes  = add(625) // 登录失败次数太多
	PasswordTooLeak     = add(628) // 密码太弱
	PasswordErr         = add(629) // 密码不正确
	TargetNumberLimit   = add(632) // 操作对象数量限制
	TargetBlocked       = add(643) // 被锁定
	UserDuplicate       = add(652) // 重复的用户
	AccessTokenExpires  = add(658) // Token 过期
	PasswordHashExpires = add(662) // 密码时间戳过期

	Degrade     = add(1200) // 被降级过滤的请求
	RPCNoClient = add(1201) // rpc服务的client都不可用
	RPCNoAuth   = add(1202) // rpc服务的client没有授权

	MobileValcodeLimitExceed       = add(10001) // 60秒下发一次验证码
	ValcodeExpires                 = add(10002) // 验证码已失效
	ValcodeWrong                   = add(10003) // 验证码错误
	AccountExist                   = add(10004) // 用户已经存在
	ShouldNotSetRefID              = add(10006) // 分类无需设置RefID
	RefIDRequired                  = add(10007) // 请输入RefID
	ChildrenIsNotAllowed           = add(10008) // 该类目不能有下级
	InvalidEmail                   = add(10009) // 邮件地址不正确
	InvalidMobile                  = add(10010) // 手机号码不正确
	InvalidGender                  = add(10011) // Gender错误
	InvalidAvatar                  = add(10012) // Avatar格式错误
	InvalidBirthYear               = add(10013) // 出生年不正确
	InvalidBirthMonth              = add(10014) // 出生月不正确
	InvalidBirthDay                = add(10015) // 出生日不正确
	SessionExpires                 = add(10016) // Session 过期
	OnlyAllowOneOwner              = add(10017) // 主理人只能有一个
	AcquireAccountIDFailed         = add(10018) // 获取用户ID失败
	MustDeleteChildrenCatalogFirst = add(10019) // 必须先删除类目子项
	NotTopicMember                 = add(10020) // 不是话题成员
	NotTopicAdmin                  = add(10021) // 不是话题管理员
	NeedIDCert                     = add(10022) // 需要身份认证
	NeedWorkCert                   = add(10023) // 需要工作认证
	OnlyAllowAdminAdded            = add(10024) // 必须管理员添加
	NeedPurchase                   = add(10025) // 必须购买
	NeedVIP                        = add(10026) // 必须是VIP用户
	NotBelongToYou                 = add(10027) // 不属于你
	NotTopicOwner                  = add(10028) // 你不是主理人
	NeedPrimaryTopic               = add(10029) // 必须有主话题
	OnlyAllowOnePrimaryTopic       = add(10030) // 只允许一个主话题
	InvalidCatalog                 = add(10031) // 不正确的类目
	NeedEditPermission             = add(10032) // 需要话题编辑权限
	ArticleEditedByOthers          = add(10033) // 文章已经被其他人修改，无法设置为私有
	GrabLinkFailed                 = add(10034) // 获取链接信息失败
	ParseHTMLFailed                = add(10035) // 解析HTML内容失败
	TopicMemberDuplicate           = add(10036) // 话题成员重复
	AuthTopicDuplicate             = add(10037) // 授权话题重复
	MustNotUseCurrentTopic         = add(10038) // 不允许授权自身
	OwnerNeedTransfer              = add(10039) // 主理人不可退出，只可转让后再退出
	SearchAccountFailed            = add(10040) // 搜索账户失败
	InviteSelfNotAllowed           = add(10041) // 不允许邀请你自己
	InvalidFeedbackType            = add(10042) // 错误的反馈类型
	AuthTopicExist                 = add(10043) // 授权话题已经存在
	MemberOverLimit                = add(10044) // 批量请求超过限制
	ModifyDiscussionNotAllowed     = add(10045) // 不能编辑该讨论
	SearchTopicFailed              = add(10046) // 搜索话题失败
	SearchArticleFailed            = add(10047) // 搜索文章失败
	SearchDiscussionFailed         = add(10048) // 搜索讨论失败
	RelFollowSelfBanned            = add(10049) // 不能关注自己
	RelFollowBlacked               = add(10050) // 被用户拉黑，无法关注
	RelFollowAlreadyBlack          = add(10051) // 已经拉黑用户，无法关注
	RelFollowAttrAlreadySet        = add(10052) // 已经设置该属性了
	RelFollowAttrNotSet            = add(10053) // 未设置该属性，不能取消
	AppExist                       = add(10054) // App 已经存在
	WorkCertExist                  = add(10055) // 工作认证已经提交
	IDCertFirst                    = add(10056) // 首选需要通过身份认证
	HasDiscussionInCategory        = add(10057) // 该分类下已经有讨论了
	NeedArticleEditPermission      = add(10058) // 需要文章编辑权限

	// 89000 - 89999 属于 Permission 类错误
	NoTopicViewPermission   = add(89001) // 没有话题查看权限
	NoTopicEditPermission   = add(89002) // 没有话题编辑权限
	NoTopicManagePermission = add(89003) // 没有话题管理权限

	// 90000 - 99999 属于 Not Exist 类错误
	UserNotExist                  = add(90001) // 用户不存在
	TagNotExist                   = add(90002) // Tag 不存在
	ConfigIdsIsEmpty              = add(90003) // ConfigIds 不存在
	ConfigsNotExist               = add(90004) // Configs 不存在
	AppNotExist                   = add(90005) // App不存在
	BuildNotExist                 = add(90006) // Build不存在
	ReviseNotExist                = add(90007) // 补充不存在
	ReviseFileNotExist            = add(90009) // 补充文件不存在
	CommentNotExist               = add(90010) // 评论不存在
	TopicInviteRequestNotExist    = add(90011) // 邀请不存在
	TopicFollowRequestNotExist    = add(90012) // 关注请求不存在
	MessageNotExist               = add(90013) // 未到找该消息
	DiscussionFileNotExist        = add(90014) // 未找到该文件
	FileNotExists                 = add(90015) // 上传文件不存在
	TopicMemberStatNotExist       = add(90016) // 话题成员统计信息出错
	DiscussionNotExist            = add(90017) // 未找到该讨论记录
	DiscussCategoryNotExist       = add(90018) // 讨论分类不存在
	AreaNotExist                  = add(90019) // 地址不存在
	LocaleNotExist                = add(90020) // 未找到该语言编码
	ArticleNotExist               = add(90021) // 文章不存在
	ArticleFileNotExist           = add(90022) // 文章附件不存在
	AccountTopicSettingNotExist   = add(90023) // 用户话题设置不存在
	RefreshTokenNotExistOrExpired = add(90024) // RefreshToken不存在
	TopicNotExist                 = add(90025) // 话题不存在
	TopicCatalogNotExist          = add(90026) // 话题类目不存在
	IDCertificationNotExist       = add(90027) // 尚未发起身份认证
	ColorNotExist                 = add(90028) // 颜色不存在
	DraftCategoryNotExist         = add(90029) // 草稿分类不存在
	DraftNotExist                 = add(90030) // 草稿不存在
	ArticleHistoryNotExist        = add(90031) // 未找到该记录
	AdminNotExist                 = add(90032) // 管理员不存在
	AccountRoleNotExist           = add(90033) // 用户角色不存在
	WorkCertificationNotExist     = add(90034) // 尚未发起工作认证

)

func IsNotExistEcode(e error) bool {
	if e == nil {
		return false
	}
	ec, ok := errors.Cause(e).(Codes)
	if ok {
		if ec.Code() > 90000 && ec.Code() < 99999 {
			return true
		}
	}

	return false
}
