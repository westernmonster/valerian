SELECT a.*
FROM topic_followers a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
