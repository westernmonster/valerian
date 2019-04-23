UPDATE oauth_access_tokens a
SET a.deleted=1
WHERE 1=1 %s
