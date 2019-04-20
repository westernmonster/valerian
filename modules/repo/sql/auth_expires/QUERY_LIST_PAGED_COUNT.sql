SELECT COUNT(1) as count
FROM auth_expires a
WHERE a.deleted=0 %s