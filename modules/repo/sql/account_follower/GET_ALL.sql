SELECT a.*
FROM account_followers a
WHERE a.deleted=0
ORDER BY a.id DESC
