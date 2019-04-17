SELECT
a.*
FROM account_followers a
WHERE a.deleted=0 %s