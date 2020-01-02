package model

const (
	MsgComment             = "comment"
	MsgReply               = "reply"
	MsgInvite              = "invite"
	MsgLike                = "like"
	MsgFollowed            = "followed"
	MsgJoined              = "joined"
	MsgApply               = "apply"
	MsgApplyRejected       = "apply_rejected"
	MsgApplyApproved       = "apply_approved"
	MsgReviseAdded         = "revise_added"
	MsgArticleCommented    = "article_commented"
	MsgReviseCommented     = "revise_commented"
	MsgDiscussionCommented = "discussion_commented"
	MsgCommentReplied      = "comment_replied"

	MsgFeedbackAccuseSuit    = "feedback_accuse_suit"
	MsgFeedbackAccuseNotSuit = "feedback_accuse_not_suit"
)

const (
	ActorTypeUser   = "user"
	ActorTypeSystem = "sytem"
)

const (
	TargetTypeTopic              = "topic"
	TargetTypeDiscussion         = "discussion"
	TargetTypeRevise             = "revise"
	TargetTypeArticle            = "article"
	TargetTypeMember             = "member"
	TargetTypeComment            = "comment"
	TargetTypeTopicInviteRequest = "invite"
	TargetTypeTopicFollowRequest = "follow"
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
	MsgTextCommentReplied      = "回复了评论"
	MsgTextFollowed            = "关注了你"
	MsgTextJoined              = "加入了话题"
	MsgTextApply               = "申请加入话题"
	MsgTextApplyRejected       = "你加入话题「%s」的申请被拒绝，原因是：%s"
	MsgTextApplyApproved       = "你已经成功加入话题「%s」"
	MsgTextInvite              = "邀请你加入话题"

	MsgTextFeedBackAccuseSuitToAuthor   = "经查您的「%s」存在违规行为，已做「%s」处理，希望您共同营造良好的社区氛围，多次违规将造成您的账号被限制"
	MsgTextFeedBackAccuseSuitToReporter = "您举报的「%s」经审核有效，处理结果「%s」，感谢您为良好的社区氛围作出贡献"
	MsgTextFeedBackAccuseNotSuit        = "您举报的「%s」经审核不存在违规信息，「%s」，感谢您为良好的社区氛围作出贡献"
)
