SELECT a.*
FROM oauth_refresh_tokens a
WHERE a.deleted=0 %s
ORDER BY a.id DESC LIMIT :limit OFFSET :offset
