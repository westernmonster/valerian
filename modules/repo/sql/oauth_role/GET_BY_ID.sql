SELECT
    a.id,
    a.name,
    a.deleted,
    a.created_at,
    a.updated_at
FROM oauth_roles a
WHERE a.id=:id AND a.deleted=0