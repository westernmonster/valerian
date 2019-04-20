SELECT COUNT(1) as count
FROM auth_access a
WHERE a.deleted=0 %s