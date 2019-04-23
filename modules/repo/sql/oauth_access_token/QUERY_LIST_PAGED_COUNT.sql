SELECT COUNT(1) as count
FROM oauth_access_tokens a
WHERE a.deleted=0 %s