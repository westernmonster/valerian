SELECT COUNT(1) as count
FROM topic_members a
WHERE a.deleted=0 %s