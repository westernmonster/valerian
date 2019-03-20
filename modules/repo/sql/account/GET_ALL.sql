SELECT a.*
FROM accounts a
WHERE a.deleted=0
ORDER BY a.id DESC
