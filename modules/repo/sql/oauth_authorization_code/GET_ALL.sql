SELECT a.*
FROM oauth_authorization_codes a
WHERE a.deleted=0
ORDER BY a.id DESC
