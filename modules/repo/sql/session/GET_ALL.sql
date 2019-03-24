SELECT a.*
FROM session a
WHERE a.deleted=0
ORDER BY a.id DESC
