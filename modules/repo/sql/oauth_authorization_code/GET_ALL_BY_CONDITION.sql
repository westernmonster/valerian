SELECT a.*
FROM oauth_authorization_codes a
WHERE a.deleted=0 %s
ORDER BY a.id DESC