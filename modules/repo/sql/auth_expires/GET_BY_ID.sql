SELECT
    a.id,
    a.token,
    a.expires_at,
    a.deleted,
    a.created_at,
    a.updated_at
FROM auth_expires a
WHERE a.id=:id AND a.deleted=0