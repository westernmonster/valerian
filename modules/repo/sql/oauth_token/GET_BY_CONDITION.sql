SELECT
a.*
FROM oauth_tokens a
WHERE a.deleted=0 %s