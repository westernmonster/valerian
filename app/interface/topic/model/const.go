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
