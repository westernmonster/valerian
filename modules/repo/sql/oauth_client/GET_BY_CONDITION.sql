SELECT
a.*
FROM oauth_clients a
WHERE a.deleted=0 %s