SELECT COUNT(1) as count
FROM auth_clients a
WHERE a.deleted=0 %s