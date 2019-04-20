UPDATE auth_expires
SET
    token=:token,
    expires_at=:expires_at,
    updated_at=:updated_at
WHERE id=:id
