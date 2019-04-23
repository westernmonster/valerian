SELECT COUNT(1) as count
FROM oauth_scopes a
WHERE a.deleted=0 %s