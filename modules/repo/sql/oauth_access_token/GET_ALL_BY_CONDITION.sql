SELECT a.*
FROM oauth_access_tokens a
WHERE a.deleted=0 %s
ORDER BY a.id DESC
