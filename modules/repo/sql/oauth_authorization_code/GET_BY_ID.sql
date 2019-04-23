SELECT
    a.id,
    a.client_id,
    a.account_id,
    a.code,
    a.redirect_uri,
    a.expires_at,
    a.scope,
    a.deleted,
    a.created_at,
    a.updated_at
FROM oauth_authorization_codes a
WHERE a.id=:id AND a.deleted=0