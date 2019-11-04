package def

const (
	PushMsgTitleArticleCommented     = "你的文章被评论了"
	PushMsgTitleDiscussionCommented  = "你的讨论被评论了"
	PushMsgTitleReviseCommented      = "你的补充被评论了"
	PushMsgTitleCommentReplied       = "你的评论被回复了"
	PushMsgTitleFollowed             = "你有了新的关注"
	PushMsgTitleArticleLiked         = "你的文章了有了新的点赞"
	PushMsgTitleReviseLiked          = "你的补充了有了新的点赞"
	PushMsgTitleDiscussionLiked      = "你的讨论了有了新的点赞"
	PushMsgTitleCommentLiked         = "你的评论了有了新的点赞"
	PushMsgTitleReviseAdded          = "你的文章有了新的补充"
	PushMsgTitleTopicFollowRequested = "你有了一条加入话题请求"
	PushMsgTitleTopicFollowApproved  = "你的加入话题请求被通过了"
	PushMsgTitleTopicFollowRejected  = "你的加入话题请求被拒绝了"
	PushMsgTitleTopicFollowInvited   = "你有一条加入话题邀请"
)

const (
	LinkArticle      = "stonote://article/%d"
	LinkRevise       = "stonote://revise/%d"
	LinkDiscussion   = "stonote://discussion/%d"
	LinkTopic        = "stonote://topic/%d"
	LinkTopicRequest = "stonote://topic_request/%d"
	LinkTopicInvite  = "stonote://topic_invite/%d"
	LinkUser         = "stonote://user/%d"
	LinkComment      = "stonote://%s/%d/comment/%d"
	LinkSubComment   = "stonote://%s/%d/comment/%d/sub/%d"
)
