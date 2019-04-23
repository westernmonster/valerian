UPDATE oauth_authorization_codes
SET
    client_id=:client_id,
    account_id=:account_id,
    code=:code,
    redirect_uri=:redirect_uri,
    expires_at=:expires_at,
    scope=:scope,
    updated_at=:updated_at
WHERE id=:id
