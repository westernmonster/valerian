SELECT COUNT(1) as count
FROM articles a
WHERE a.deleted=0 %s