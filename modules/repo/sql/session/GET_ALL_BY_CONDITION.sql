SELECT a.*
FROM session a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
