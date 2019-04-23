SELECT
    a.id,
    a.client_id,
    a.account_id,
    a.token,
    a.expires_at,
    a.scope,
    a.deleted,
    a.created_at,
    a.updated_at
FROM oauth_refresh_tokens a
WHERE a.id=:id AND a.deleted=0