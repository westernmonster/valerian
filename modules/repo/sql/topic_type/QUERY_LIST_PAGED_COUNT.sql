SELECT COUNT(1) as count
FROM topic_types a
WHERE a.deleted=0 %s