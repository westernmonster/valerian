SELECT a.*
FROM oauth_access_tokens a
WHERE a.deleted=0
ORDER BY a.id DESC
