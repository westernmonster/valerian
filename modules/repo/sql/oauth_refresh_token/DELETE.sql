UPDATE oauth_refresh_tokens
SET deleted=1
WHERE id=:id
