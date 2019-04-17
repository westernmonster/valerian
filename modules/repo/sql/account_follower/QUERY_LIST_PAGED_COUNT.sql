SELECT COUNT(1) as count
FROM account_followers a
WHERE a.deleted=0 %s