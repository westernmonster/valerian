package model

const (
	TopicCatalogTaxonomy = "taxonomy"
	TopicCatalogArticle  = "article"
	TopicCatalogTestSet  = "test_set"
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
	MemberRoleOwner = "owner"
	MemberRoleAdmin = "admin"
	MemberRoleUser  = "user"
)

const (
	TopicRelationStrong = "strong"
	TopicRelationNormal = "normal"
)

const (
	FollowStatusUnfollowed = 0
	FollowStatusApproving  = 1
	FollowStatusFollowed   = 2
)

const (
	FollowRequestStatusCommited = 0
	FollowRequestStatusApproved = 1
	FollowRequestStatusRejected = 2
)

const (
	ReportTypeSpam       = 1 // 垃圾广告
	ReportTypeCopyRight  = 2 // 涉嫌侵权
	ReportTypeDiscomfort = 3 // 内容引起不适
	ReportTypeIncorrect  = 4 // 内容有误
	ReportTypeOther      = 5 // 其他
)
