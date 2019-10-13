package model

const (
	MsgComment       = "comment"
	MsgReply         = "reply"
	MsgInvite        = "invite"
	MsgApply         = "apply"
	MsgLike          = "like"
	MsgFollowed      = "followed"
	MsgJoined        = "joined"
	MsgApplyRejected = "joined"
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
)

const (
	MsgTextLikeArticle    = "赞了文章"
	MsgTextLikeRevise     = "赞了补充"
	MsgTextLikeComment    = "赞了评论"
	MsgTextLikeDiscussion = "赞了讨论"
)
