UPDATE oauth_scopes
SET
    scope=:scope,
    description=:description,
    is_default=:is_default,
    updated_at=:updated_at
WHERE id=:id
