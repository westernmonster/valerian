SELECT COUNT(1) as count
FROM oauth_clients a
WHERE a.deleted=0 %s