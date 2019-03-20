SELECT a.*
FROM accounts a
WHERE a.deleted=0 %s
ORDER BY a.id DESC LIMIT :limit OFFSET :offset
