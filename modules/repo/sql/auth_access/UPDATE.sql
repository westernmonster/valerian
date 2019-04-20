UPDATE auth_access
SET
    client_id=:client_id,
    authorize=:authorize,
    previous=:previous,
    access_token=:access_token,
    refresh_token=:refresh_token,
    expired_in=:expired_in,
    scope=:scope,
    redirect_uri=:redirect_uri,
    extra=:extra,
    updated_at=:updated_at
WHERE id=:id
