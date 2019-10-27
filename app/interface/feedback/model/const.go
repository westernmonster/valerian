package model

const (
	TargetTypeFeedback = int32(1)
	TargetTypeMember   = int32(2)
	TargetTypeTopic    = int32(3)
	TargetTypeArticle  = int32(4)
	TargetTypeDiscuss  = int32(5)
	TargetTypeRevise   = int32(6)
	TargetTypeComment  = int32(7)
)

const (
	FeedbackTypeAccusePeople  = "accuse_people"
	FeedbackTypeAccuseContent = "accuse_content"
	FeedbackTypeFeedback      = "feedback"
)
