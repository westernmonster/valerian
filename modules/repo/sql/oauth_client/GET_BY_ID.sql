SELECT
    a.id,
    a.client_id,
    a.client_secret,
    a.extra,
    a.redirect_uri,
    a.deleted,
    a.created_at,
    a.updated_at
FROM oauth_clients a
WHERE a.id=:id AND a.deleted=0