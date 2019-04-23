SELECT a.*
FROM oauth_scopes a
WHERE a.deleted=0
ORDER BY a.id DESC
