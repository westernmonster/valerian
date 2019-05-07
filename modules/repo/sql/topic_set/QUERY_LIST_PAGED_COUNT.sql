SELECT COUNT(1) as count
FROM topic_sets a
WHERE a.deleted=0 %s