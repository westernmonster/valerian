UPDATE oauth_access_tokens
SET
    client_id=:client_id,
    account_id=:account_id,
    token=:token,
    expires_at=:expires_at,
    scope=:scope,
    updated_at=:updated_at
WHERE id=:id
