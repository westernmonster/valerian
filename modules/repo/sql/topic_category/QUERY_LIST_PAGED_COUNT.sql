SELECT COUNT(1) as count
FROM topic_categories a
WHERE a.deleted=0 %s