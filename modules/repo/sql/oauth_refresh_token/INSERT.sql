INSERT INTO oauth_refresh_tokens(
    id,
    client_id,
    account_id,
    token,
    expires_at,
    scope,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :client_id,
    :account_id,
    :token,
    :expires_at,
    :scope,
    :deleted,
    :created_at,
    :updated_at
)