SELECT a.*
FROM auth_access a
WHERE a.deleted=0 %s
ORDER BY a.id DESC LIMIT :limit OFFSET :offset
