SELECT a.*
FROM auth_expires a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
