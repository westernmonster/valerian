INSERT INTO auth_expires(
    id,
    token,
    expires_at,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :token,
    :expires_at,
    :deleted,
    :created_at,
    :updated_at
)