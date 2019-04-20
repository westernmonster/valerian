UPDATE auth_authorize
SET
    client_id=:client_id,
    code=:code,
    expired_in=:expired_in,
    scope=:scope,
    redirect_uri=:redirect_uri,
    state=:state,
    extra=:extra,
    updated_at=:updated_at
WHERE id=:id
