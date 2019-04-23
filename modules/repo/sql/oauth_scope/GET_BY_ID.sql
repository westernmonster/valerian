SELECT
    a.id,
    a.scope,
    a.description,
    a.is_default,
    a.deleted,
    a.created_at,
    a.updated_at
FROM oauth_scopes a
WHERE a.id=:id AND a.deleted=0