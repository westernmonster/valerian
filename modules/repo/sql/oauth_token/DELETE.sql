UPDATE oauth_tokens
SET deleted=1
WHERE id=:id
