SELECT a.*
FROM oauth_scopes a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
