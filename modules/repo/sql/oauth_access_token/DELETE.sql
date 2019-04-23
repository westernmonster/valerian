UPDATE oauth_access_tokens
SET deleted=1
WHERE id=:id
