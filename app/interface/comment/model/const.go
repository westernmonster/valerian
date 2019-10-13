package model

const (
	TargetTypeTopic      = "topic"
	TargetTypeDiscussion = "discussion"
	TargetTypeRevise     = "revise"
	TargetTypeArticle    = "article"
	TargetTypeMember     = "member"
	TargetTypeComment    = "comment"
)

func IsValidTargetType(ctype string) bool {
	switch ctype {
	case
		TargetTypeRevise,
		TargetTypeArticle,
		TargetTypeComment,
		TargetTypeDiscussion:
		return true
	}
	return false
}
