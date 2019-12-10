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
	// 未关注
	FollowStatusUnfollowed = 1
	// 审核中
	FollowStatusApproving = 2
	// 已关注
	FollowStatusFollowed = 3
)

const (
	InviteStatusSent     = 1
	InviteStatusJoined   = 2
	InviteStatusRejected = 3
)

const (
	// 审核请求已经提交
	FollowRequestStatusCommited = 0
	// 审核请求通过
	FollowRequestStatusApproved = 1
	// 审核请求拒绝
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
	TargetTypeMember         = "member"
)
