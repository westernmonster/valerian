SELECT a.*
FROM oauth_roles a
WHERE a.deleted=0
ORDER BY a.id DESC
