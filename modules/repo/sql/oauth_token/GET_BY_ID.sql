SELECT
    a.id,
    a.expired_at,
    a.code,
    a.access,
    a.refresh,
    a.data,
    a.deleted,
    a.created_at,
    a.updated_at
FROM oauth_tokens a
WHERE a.id=:id AND a.deleted=0