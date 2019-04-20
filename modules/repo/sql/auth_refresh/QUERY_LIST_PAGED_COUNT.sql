SELECT COUNT(1) as count
FROM auth_refresh a
WHERE a.deleted=0 %s