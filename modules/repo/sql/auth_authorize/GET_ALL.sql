SELECT a.*
FROM auth_authorize a
WHERE a.deleted=0
ORDER BY a.id DESC
