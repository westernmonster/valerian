SELECT
    a.id,
    a.client_id,
    a.authorize,
    a.previous,
    a.access_token,
    a.refresh_token,
    a.expired_in,
    a.scope,
    a.redirect_uri,
    a.extra,
    a.deleted,
    a.created_at,
    a.updated_at
FROM auth_access a
WHERE a.id=:id AND a.deleted=0