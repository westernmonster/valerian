INSERT INTO auth_refresh(
    id,
    token,
    access,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :token,
    :access,
    :deleted,
    :created_at,
    :updated_at
)