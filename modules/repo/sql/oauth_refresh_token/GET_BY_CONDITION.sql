SELECT
a.*
FROM oauth_refresh_tokens a
WHERE a.deleted=0 %s