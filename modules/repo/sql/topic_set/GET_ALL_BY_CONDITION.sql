SELECT a.*
FROM topic_sets a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
