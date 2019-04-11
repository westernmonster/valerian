SELECT a.*
FROM topic_members a
WHERE a.deleted=0
ORDER BY a.id DESC
