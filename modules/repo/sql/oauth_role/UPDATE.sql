UPDATE oauth_roles
SET
    name=:name,
    updated_at=:updated_at
WHERE id=:id
