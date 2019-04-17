SELECT COUNT(1) as count
FROM topic_followers a
WHERE a.deleted=0 %s