SELECT COUNT(1) as count
FROM oauth_authorization_codes a
WHERE a.deleted=0 %s