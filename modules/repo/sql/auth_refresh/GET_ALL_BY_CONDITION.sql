SELECT a.*
FROM auth_refresh a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
