SELECT COUNT(1) as count
FROM auth_authorize a
WHERE a.deleted=0 %s