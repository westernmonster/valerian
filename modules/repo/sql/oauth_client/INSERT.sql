INSERT INTO oauth_clients(
    id,
    client_id,
    client_secret,
    extra,
    redirect_uri,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :client_id,
    :client_secret,
    :extra,
    :redirect_uri,
    :deleted,
    :created_at,
    :updated_at
)