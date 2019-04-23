SELECT COUNT(1) as count
FROM oauth_roles a
WHERE a.deleted=0 %s