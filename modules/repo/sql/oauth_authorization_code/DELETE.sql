UPDATE oauth_authorization_codes
SET deleted=1
WHERE id=:id
