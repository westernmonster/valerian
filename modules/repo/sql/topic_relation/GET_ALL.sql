SELECT a.*
FROM topic_relations a
WHERE a.deleted=0
ORDER BY a.id DESC
