INSERT INTO oauth_scopes(
    id,
    scope,
    description,
    is_default,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :scope,
    :description,
    :is_default,
    :deleted,
    :created_at,
    :updated_at
)