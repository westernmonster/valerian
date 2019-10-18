package ecode

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
	FileNotExists       = add(616) // 上传文件不存在
	FileTooLarge        = add(617) // 上传文件太大
	FailedTooManyTimes  = add(625) // 登录失败次数太多
	UserNotExist        = add(626) // 用户不存在
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
	AreaNotExist                   = add(10005) // 地址不存在
	ShouldNotSetRefID              = add(10006) // 分类无需设置RefID
	RefIDRequired                  = add(10007) // 请输入RefID
	ChildrenIsNotAllowed           = add(10008) // 该类目不能有下级
	InvalidEmail                   = add(10009) // 邮件地址不正确
	InvalidMobile                  = add(11010) // 手机号码不正确
	InvalidGender                  = add(11011) // Gender错误
	InvalidAvatar                  = add(11012) // Avatar格式错误
	InvalidBirthYear               = add(11013) // 出生年不正确
	InvalidBirthMonth              = add(11014) // 出生月不正确
	InvalidBirthDay                = add(11015) // 出生日不正确
	SessionExpires                 = add(11016) // Session 过期
	RefreshTokenNotExistOrExpired  = add(11017) // RefreshToken不存在
	TopicNotExist                  = add(11018) // 话题不存在
	OnlyAllowOneOwner              = add(11019) // 主理人只能有一个
	AcquireAccountIDFailed         = add(11020) // 获取用户ID失败
	TopicCatalogNotExist           = add(11032) // 话题类目不存在
	MustDeleteChildrenCatalogFirst = add(11033) // 必须先删除类目子项
	NotTopicMember                 = add(11034) // 不是话题成员
	NotTopicAdmin                  = add(11035) // 不是话题管理员
	NeedIDCert                     = add(11037) // 需要身份认证
	NeedWorkCert                   = add(11038) // 需要工作认证
	OnlyAllowAdminAdded            = add(11039) // 必须管理员添加
	NeedPurchase                   = add(11040) // 必须购买
	NeedVIP                        = add(11041) // 必须是VIP用户
	IDCertificationNotExist        = add(11042) // 尚未发起身份认证
	ColorNotExist                  = add(11043) // 颜色不存在
	NotBelongToYou                 = add(11044) // 不属于你
	DraftCategoryNotExist          = add(11045) // 草稿分类不存在
	NotTopicOwner                  = add(11046) // 你不是主理人
	DraftNotExist                  = add(11047) // 草稿不存在
	NeedPrimaryTopic               = add(11049) // 必须有主话题
	OnlyAllowOnePrimaryTopic       = add(11050) // 只允许一个主话题
	InvalidCatalog                 = add(11052) // 不正确的类目
	NeedEditPermission             = add(11054) // 需要话题编辑权限
	LocaleNotExist                 = add(11055) // 未找到该语言编码
	ArticleEditedByOthers          = add(11056) // 文章已经被其他人修改，无法设置为私有
	ArticleNotExist                = add(11057) // 文章不存在
	ArticleFileNotExist            = add(11058) // 文章附件不存在
	AccountTopicSettingNotExist    = add(11059) // 用户话题设置不存在
	GrabLinkFailed                 = add(11063) // 获取链接信息失败
	ParseHTMLFailed                = add(11064) // 解析HTML内容失败
	TopicMemberDuplicate           = add(11065) // 话题成员重复
	AuthTopicDuplicate             = add(11066) // 授权话题重复
	MustNotUseCurrentTopic         = add(11067) // 不允许授权自身
	OwnerNeedTransfer              = add(11068) // 主理人不可退出，只可转让后再退出
	DiscussCategoryNotExist        = add(11069) // 讨论分类不存在
	SearchAccountFailed            = add(11070) // 搜索账户失败
	InviteSelfNotAllowed           = add(11071) // 不允许邀请你自己
	InvalidFeedbackType            = add(11072) // 错误的反馈类型
	AuthTopicExist                 = add(11073) // 授权话题已经存在
	ArticleHistoryNotExist         = add(11074) // 未找到该记录
	MemberOverLimit                = New(11075) // 批量请求超过限制
	TopicMemberStatNotExist        = New(10078) // 话题成员统计信息出错
	DiscussionNotExist             = New(10079) // 未找到该讨论记录
	ModifyDiscussionNotAllowed     = New(10080) // 不能编辑该讨论
	DiscussionFileNotExist         = New(10081) // 未找到该文件
	ReviseNotExist                 = New(10082) // 补充不存在
	ReviseFileNotExist             = New(10083) // 补充文件不存在
	CommentNotExist                = New(10084) // 评论不存在
	TopicInviteRequestNotExist     = New(10085) // 邀请不存在
	TopicFollowRequestNotExist     = New(10086) // 关注请求不存在

	TagNotExist      = add(20000) // Tag 不存在
	ConfigIdsIsEmpty = add(20001) // ConfigIds 不存在
	ConfigsNotExist  = add(20002) // Configs 不存在
	AppNotExist      = add(20003) // App不存在
	BuildNotExist    = add(20004) // Build不存在

	// Relation
	RelFollowSelfBanned     = New(20005) // 不能关注自己
	RelFollowBlacked        = New(20006) // 被用户拉黑，无法关注
	RelFollowAlreadyBlack   = New(20007) // 已经拉黑用户，无法关注
	RelFollowAttrAlreadySet = New(20008) // 已经设置该属性了
	RelFollowAttrNotSet     = New(20009) // 未设置该属性，不能取消

)
