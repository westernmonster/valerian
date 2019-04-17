SELECT a.*
FROM topic_followers a
WHERE a.deleted=0
ORDER BY a.id DESC
