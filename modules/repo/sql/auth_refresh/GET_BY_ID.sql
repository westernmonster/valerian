SELECT
    a.id,
    a.token,
    a.access,
    a.deleted,
    a.created_at,
    a.updated_at
FROM auth_refresh a
WHERE a.id=:id AND a.deleted=0