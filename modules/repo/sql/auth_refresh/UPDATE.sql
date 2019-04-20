UPDATE auth_refresh
SET
    token=:token,
    access=:access,
    updated_at=:updated_at
WHERE id=:id
