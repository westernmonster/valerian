package models

import (
	"time"
)

const (
	DefaultAvatarUrl = "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png"
)

const (
	ValcodeSpan = time.Minute
)

const (
	JWTSignKey = "ETEcP/nuY6YQT5KM8xWb2+04QG4nbek7H0sowM42dihER/e2giDFBpDSI5pGIDTn2C4SLxn0WP+qW3ZlUUSXLg=="
	TokenKey   = "flywiki_token"

	OAUTH2MobileClientID     = "532c28d5412dd75bf975fb951c740a30"
	OAUTH2MobileClientSecret = "16ed0e1ab220aebf9362045ccad0664f"
	OAUTH2MobileClientDomain = "https://api.flywk.com"

	OAUTH2WebClientID     = "2567a5ec9705eb7ac2c984033e06189d"
	OAUTH2WebClientSecret = "8b17d5515cdc1939d83abe5c00d673ad"
	OAUTH2WebClientDomain = "https://www.flywk.com"

	GrantTypePassword     = "password"
	GrantTypeDigits       = "digits"
	GrantTypeRefreshToken = "refresh_token"
	GrantTypeGoogle       = "google"
	GrantTypeWeChat       = "wechat"

	ResponseTypeToken = "token"
	ResponseTypeCode  = "code"

	Issuer = "https://flywk.com"

	PasswordPepper = "8PrhfDL7Qr/G6%RbaiQouFfQ"

	ExpiresIn = 60 * 60 * 72

	// GrantTypeSina         = "sina"
	// GrantTypeQQConn       = "qqconn"
	// GrantTypeCTCC         = "ctcc"
	// GrantTypeCMCC         = "cmcc"
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
	SourceWeb     = 1
	SourceiOS     = 2
	SourceAndroid = 3
)

const (
	// 手机号码正则表达式
	// Refer: https://github.com/VincentSit/ChinaMobilePhoneNumberRegex
	// Phone number + Data only number)
	ChinaMobileRegex = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`

	OtherMobileRegex = `^\d+$`
)

const (
	SessionTypeResetPassword = 1
)

const (
	GenderMale   = 1
	GenderFemale = 2
)

const (
	IDCertificationUncommitted = -1
	IDCertificationInProgress  = 0
	IDCertificationSuccess     = 1
	IDCertificationFailed      = 2
)

const (
	HeaderStatusCode = "flywiki-status-code"
)

const (
	JoinPermissionMember          = "member"
	JoinPermissionIDCert          = "id_cert"
	JoinPermissionWorkCert        = "work_cert"
	JoinPermissionMemberApprove   = "member_approve"
	JoinPermissionIDCertApprove   = "id_cert_approve"
	JoinPermissionWorkCertApprove = "work_cert_approve"
	JoinPermissionAdminAdd        = "admin_add"
	JoinPermissionPurchase        = "purchase"
	JoinPermissionVIP             = "vip"
)

const (
	ViewPermissionPublic = "public"
	ViewPermissionJoin   = "join"
)

const (
	EditPermissionIDCert                 = "id_cert"
	EditPermissionWorkCert               = "work_cert"
	EditPermissionIDCertJoined           = "id_cert_joined"
	EditPermissionWorkCertJoined         = "work_cert_joined"
	EditPermissionApprovedIDCertJoined   = "approved_id_cert_joined"
	EditPermissionApprovedWorkCertJoined = "approved_work_cert_joined"
	EditPermissionAdmin                  = "only_admin"
)

const (
	CatalogViewTypeSection = "section"
	CatalogViewTypeColumn  = "column"
)

const (
	TopicHomeIntroduction = "introduction"
	TopicHomeFeed         = "feed"
	TopicHomeCataglog     = "catalog"
	TopicHomeDiscussion   = "discussion"
	TopicHomeChat         = "chat"
)

const (
	AccountRoleAdmin = "admin"
	AccountRoleUser  = "user"
	AccountRoleOrg   = "org"
)

const (
	MemberRoleOwner = "owner"
	MemberRoleAdmin = "admin"
	MemberRoleUser  = "user"
)

const (
	TopicRelationStrong = "strong"
	TopicRelationNormal = "normal"
)

const (
	TopicCatalogTaxonomy = "taxonomy"
	TopicCatalogArticle  = "article"
	TopicCatalogTestSet  = "test_set"
)

// x-app-version App 版本号  3.22
// x-app-versioncode App 版本号 1024
// x-app-build release/debug
// x-app-za OS=iOS&Release=12.2&Model=iPhone7,2&VersionName=4.39.0&VersionCode=1280&Width=750&Height=1334&DeviceType=Phone&Brand=Apple&OperatorType=46011
// x-network-type 4G/3G etc
// x-udid 设备唯一标识符
