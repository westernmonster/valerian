UPDATE oauth_tokens
SET
    expired_at=:expired_at,
    code=:code,
    access=:access,
    refresh=:refresh,
    data=:data,
    updated_at=:updated_at
WHERE id=:id
