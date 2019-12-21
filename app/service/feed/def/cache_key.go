package def

import fmt "fmt"

func RefreshTokenKey(token string) string {
	return fmt.Sprintf("ak_%s", token)
}

func MobileValcodeKey(vtype int32, mobile string) string {
	return fmt.Sprintf("rc_%d_%s", vtype, mobile)
}

func EmailValcodeKey(vtype int32, email string) string {
	return fmt.Sprintf("rc_%d_%s", vtype, email)
}

func LocaleKey() string {
	return "locales"
}

func WorkcertKey(aid int64) string {
	return fmt.Sprintf("workcert_%d", aid)
}

func AccountKey(aid int64) string {
	return fmt.Sprintf("account_%d", aid)
}

func IdcertKey(aid int64) string {
	return fmt.Sprintf("idcert_%d", aid)
}

func ArticleKey(articleID int64) string {
	return fmt.Sprintf("article_%d", articleID)
}

func ReviseFileKey(reviseID int64) string {
	return fmt.Sprintf("rev_files_%d", reviseID)
}

func ArticleFileKey(articleID int64) string {
	return fmt.Sprintf("a_files_%d", articleID)
}

func ArticleHistoryKey(articleVersionID int64) string {
	return fmt.Sprintf("a_history_%d", articleVersionID)
}

func TopicCatalogKey(topicID int64) string {
	return fmt.Sprintf("t_catalog_%d", topicID)
}

func ReviseKey(reviseID int64) string {
	return fmt.Sprintf("revise_%d", reviseID)
}

func FansKey(aid int64, page, pageSize int, version string) string {
	return fmt.Sprintf("fans_%d_%d_%d_%s", aid, page, pageSize, version)
}

func FansVersionKey(aid int64) string {
	return fmt.Sprintf("fansv_%d", aid)
}

func FollowingsKey(aid int64, page, pageSize int, version string) string {
	return fmt.Sprintf("fll_%d_%d_%d_%s", aid, page, pageSize, version)
}

func FollowingVersionKey(aid int64) string {
	return fmt.Sprintf("fllv_%d", aid)
}

func TopicKey(topicID int64) string {
	return fmt.Sprintf("t_%d", topicID)
}

func AccountTopicSettingKey(aid int64, topicID int64) string {
	return fmt.Sprintf("acc_topic_setting_%d_%d", aid, topicID)
}

func AuthTopicsKey(topicID int64) string {
	return fmt.Sprintf("auth_topics_%d", topicID)
}

func TopicMembersKey(topicID int64, page, pageSize int32, version string) string {
	return fmt.Sprintf("tms_%d_%d_%d_%s", topicID, page, pageSize, version)
}

func TopicMemberVersionKey(topicID int64) string {
	return fmt.Sprintf("tmv_%d", topicID)
}

func DiscussionKey(discussionID int64) string {
	return fmt.Sprintf("d_%d", discussionID)
}

func DiscussionFileKey(discussionID int64) string {
	return fmt.Sprintf("d_files_%d", discussionID)
}

func DiscussionCategoriesKey(topicID int64) string {
	return fmt.Sprintf("d_cates_%d", topicID)
}

func ResetPasswordKey(sessionID string) string {
	return fmt.Sprintf("srp_%s", sessionID)
}

func SessionKey(sid string) string {
	return fmt.Sprintf("sess_%d", sid)
}
