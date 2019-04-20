SELECT a.*
FROM auth_refresh a
WHERE a.deleted=0
ORDER BY a.id DESC
