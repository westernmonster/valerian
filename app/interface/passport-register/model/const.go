package model

const (
	PasswordPepper = "8PrhfDL7Qr/G6%RbaiQouFfQ"
)

const (
	SourceWeb     = int32(1)
	SourceiOS     = int32(2)
	SourceAndroid = int32(3)
)

const (
	// 手机号码正则表达式
	// Refer: https://github.com/VincentSit/ChinaMobilePhoneNumberRegex
	// Phone number + Data only number)
	ChinaMobileRegex = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`

	OtherMobileRegex = `^\d+$`
)

const (
	ValcodeRegister       = int32(1)
	ValcodeForgetPassword = int32(2)
	ValcodeLogin          = int32(3)
)

const (
	AccountRoleUser       = "user"
	AccountRoleAdmin      = "admin"
	AccountRoleSuperAdmin = "superadmin"
	AccountRoleOrg        = "org"
)

const (
	WebCookieExpires = 60 * 60 * 72
)
