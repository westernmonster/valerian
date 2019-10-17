package model

const (
	MsgComment             = "comment"
	MsgReply               = "reply"
	MsgInvite              = "invite"
	MsgApply               = "apply"
	MsgLike                = "like"
	MsgFollowed            = "followed"
	MsgJoined              = "joined"
	MsgApplyRejected       = "apply_rejected"
	MsgReviseAdded         = "revise_added"
	MsgArticleCommented    = "article_commented"
	MsgReviseCommented     = "revise_commented"
	MsgDiscussionCommented = "discussion_commented"
)

const (
	ActorTypeUser   = "user"
	ActorTypeSystem = "sytem"
)

const (
	TargetTypeTopic      = "topic"
	TargetTypeDiscussion = "discussion"
	TargetTypeRevise     = "revise"
	TargetTypeArticle    = "article"
	TargetTypeMember     = "member"
	TargetTypeComment    = "comment"
)

const (
	MsgTextLikeArticle         = "赞了文章"
	MsgTextLikeRevise          = "赞了补充"
	MsgTextLikeComment         = "赞了评论"
	MsgTextLikeDiscussion      = "赞了讨论"
	MsgTextReviseAdded         = "补充了文章"
	MsgTextArticleCommented    = "评论了文章"
	MsgTextReviseCommented     = "评论了补充"
	MsgTextDiscussionCommented = "评论了讨论"
)
