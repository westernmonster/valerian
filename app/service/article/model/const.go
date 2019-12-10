package model

const (
	TopicCatalogTaxonomy = "taxonomy"
	TopicCatalogArticle  = "article"
	TopicCatalogTestSet  = "test_set"
	TopicCatalogTopic    = "topic"
)

const (
	JoinPermissionMember        = "member"
	JoinPermissionMemberApprove = "member_approve"
	JoinPermissionCertApprove   = "cert_approve"
	JoinPermissionManualAdd     = "manual_add"
)

const (
	ViewPermissionPublic = "public"
	ViewPermissionJoin   = "join"
)

const (
	EditPermissionMember = "member"
	EditPermissionAdmin  = "admin"
)

const (
	CatalogViewTypeSection = "section"
	CatalogViewTypeColumn  = "column"
)

const (
	TopicHomeFeed     = "feed"
	TopicHomeCataglog = "catalog"
	TopicHomeDiscuss  = "discuss"
)

const (
	MemberRoleOwner = "owner"
	MemberRoleAdmin = "admin"
	MemberRoleUser  = "user"
)

const (
	AuthPermissionView      = "view"
	AuthPermissionEdit      = "edit"
	AuthPermissionAdminEdit = "admin_edit"
)

const (
	FollowStatusUnfollowed = 0
	FollowStatusApproving  = 1
	FollowStatusFollowed   = 2
)

const (
	InviteStatusSent = 1
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

const (
	TargetTypeTopic          = "topic"
	TargetTypeDiscussion     = "discussion"
	TargetTypeRevise         = "revise"
	TargetTypeArticle        = "article"
	TargetTypeArticleHistory = "article_history"
	TargetTypeComment        = "comment"
)

const (
	UserRoleAdmin      = "admin"
	UserRoleSuperAdmin = "superadmin"
	UserRoleUser       = "user"
	UserRoleOrg        = "org"
)

const (
	FileTypeWord  = "word"
	FileTypePPT   = "ppt"
	FileTypeExcel = "excel"
	FileTypePDF   = "pdf"
)
