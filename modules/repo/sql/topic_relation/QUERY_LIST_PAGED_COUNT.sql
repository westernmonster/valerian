SELECT COUNT(1) as count
FROM topic_relations a
WHERE a.deleted=0 %s