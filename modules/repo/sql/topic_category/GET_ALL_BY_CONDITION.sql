SELECT a.*
FROM topic_categories a
WHERE a.deleted=0 %s
ORDER BY a.id DESC