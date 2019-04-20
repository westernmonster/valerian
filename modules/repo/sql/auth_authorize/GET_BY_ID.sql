SELECT
    a.id,
    a.client_id,
    a.code,
    a.expired_in,
    a.scope,
    a.redirect_uri,
    a.state,
    a.extra,
    a.deleted,
    a.created_at,
    a.updated_at
FROM auth_authorize a
WHERE a.id=:id AND a.deleted=0