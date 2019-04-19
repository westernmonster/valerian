INSERT INTO oauth_tokens(
    id,
    expired_at,
    code,
    access,
    refresh,
    data,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :expired_at,
    :code,
    :access,
    :refresh,
    :data,
    :deleted,
    :created_at,
    :updated_at
)