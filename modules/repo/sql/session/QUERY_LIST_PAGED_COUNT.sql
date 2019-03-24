SELECT COUNT(1) as count
FROM session a
WHERE a.deleted=0 %s