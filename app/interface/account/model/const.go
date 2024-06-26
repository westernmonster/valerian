package model

const (
	IdentityMobile = int32(1)
	IdentityEmail  = int32(2)
)

const (
	// 手机号码正则表达式
	// Refer: https://github.com/VincentSit/ChinaMobilePhoneNumberRegex
	// Phone number + Data only number)
	ChinaMobileRegex = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`

	OtherMobileRegex = `^\d+$`
)

const (
	PasswordPepper = "8PrhfDL7Qr/G6%RbaiQouFfQ"
)

const (
	GenderMale   = int32(1)
	GenderFemale = int32(2)
)

const (
	ValcodeRegister       = int32(1)
	ValcodeForgetPassword = int32(2)
	ValcodeLogin          = int32(3)
	ValcodeClose          = int32(4)
)

const (
	TargetTypeTopic          = "topic"
	TargetTypeDiscussion     = "discussion"
	TargetTypeRevise         = "revise"
	TargetTypeArticle        = "article"
	TargetTypeMember         = "member"
	TargetTypeComment        = "comment"
	TargetTypeArticleHistory = "article_history"
)

const (
	CateManaged      = "managed"
	CateFollowed     = "followed"
	CateFaved        = "faved"
	CateRecentViewed = "viewed"
	CateCreated      = "created"
)

const (
	IDCertificationUncommitted = int32(-1)
	IDCertificationInProgress  = int32(0)
	IDCertificationSuccess     = int32(1)
	IDCertificationFailed      = int32(2)
)

const (
	WorkCertificationUncommitted = int32(-1)
	WorkCertificationInProgress  = int32(0)
	WorkCertificationSuccess     = int32(1)
	WorkCertificationFailed      = int32(2)
)
