package model

const (
	PasswordPepper = "8PrhfDL7Qr/G6%RbaiQouFfQ"
)

const (
	SourceWeb     = 1
	SourceiOS     = 2
	SourceAndroid = 3
)

const (
	ResponseTypeToken = "token"
	ResponseTypeCode  = "code"
)

const (
	// 手机号码正则表达式
	// Refer: https://github.com/VincentSit/ChinaMobilePhoneNumberRegex
	// Phone number + Data only number)
	ChinaMobileRegex = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`

	OtherMobileRegex = `^\d+$`
)

const (
	IdentityMobile = 1
	IdentityEmail  = 2
)

const (
	ValcodeRegister       = 1
	ValcodeForgetPassword = 2
	ValcodeLogin          = 3
)

const (
	WebCookieExpires = 60 * 60 * 72
)
