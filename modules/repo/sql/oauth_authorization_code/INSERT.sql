INSERT INTO oauth_authorization_codes(
    id,
    client_id,
    account_id,
    code,
    redirect_uri,
    expires_at,
    scope,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :client_id,
    :account_id,
    :code,
    :redirect_uri,
    :expires_at,
    :scope,
    :deleted,
    :created_at,
    :updated_at
)