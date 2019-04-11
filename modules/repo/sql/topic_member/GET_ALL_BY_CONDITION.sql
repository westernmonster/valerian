SELECT a.*
FROM topic_members a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
