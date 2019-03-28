SELECT a.*
FROM topics a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
