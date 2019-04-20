INSERT INTO auth_authorize(
    id,
    client_id,
    code,
    expired_in,
    scope,
    redirect_uri,
    state,
    extra,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :client_id,
    :code,
    :expired_in,
    :scope,
    :redirect_uri,
    :state,
    :extra,
    :deleted,
    :created_at,
    :updated_at
)