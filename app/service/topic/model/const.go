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
	FollowStatusUnfollowed = int32(1)
	// 审核中
	FollowStatusApproving = int32(2)
	// 已关注
	FollowStatusFollowed = int32(3)
)

const (
	InviteStatusSent     = int32(1)
	InviteStatusJoined   = int32(2)
	InviteStatusRejected = int32(3)
)

const (
	// 审核请求已经提交
	FollowRequestStatusCommited = int32(0)
	// 审核请求通过
	FollowRequestStatusApproved = int32(1)
	// 审核请求拒绝
	FollowRequestStatusRejected = int32(2)
)

const (
	ReportTypeSpam       = int32(1) // 垃圾广告
	ReportTypeCopyRight  = int32(2) // 涉嫌侵权
	ReportTypeDiscomfort = int32(3) // 内容引起不适
	ReportTypeIncorrect  = int32(4) // 内容有误
	ReportTypeOther      = int32(5) // 其他
)

const (
	TargetTypeTopic      = "topic"
	TargetTypeDiscussion = "discussion"
	TargetTypeRevise     = "revise"
	TargetTypeArticle    = "article"
	TargetTypeMember     = "member"
)
