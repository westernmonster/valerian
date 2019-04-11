SELECT a.*
FROM articles a
WHERE a.deleted=0
ORDER BY a.id DESC
