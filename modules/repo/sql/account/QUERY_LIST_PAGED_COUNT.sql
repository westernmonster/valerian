SELECT COUNT(1) as count
FROM accounts a
WHERE a.deleted=0 %s