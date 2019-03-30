package models

import (
	"time"
)

const (
	ValcodeSpan = time.Minute
)

const (
	ValcodeRegister       = 1
	ValcodeForgetPassword = 2
)

const (
	SourceWeb     = 1
	SourceiOS     = 2
	SourceAndroid = 3
)

const (
	// 手机号码正则表达式
	// Refer: https://github.com/VincentSit/ChinaMobilePhoneNumberRegex
	// Phone number + Data only number)
	MobileRegex = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`
)

const (
	SessionTypeResetPassword = 1
)

const (
	GenderMale   = 1
	GenderFemale = 2
)
