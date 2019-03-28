SELECT COUNT(1) as count
FROM topics a
WHERE a.deleted=0 %s