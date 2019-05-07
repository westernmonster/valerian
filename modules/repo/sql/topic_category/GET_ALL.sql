SELECT a.*
FROM topic_categories a
WHERE a.deleted=0
ORDER BY a.id DESC
