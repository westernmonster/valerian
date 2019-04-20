SELECT a.*
FROM auth_access a
WHERE a.deleted=0
ORDER BY a.id DESC
