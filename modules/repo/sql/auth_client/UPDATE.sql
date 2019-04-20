UPDATE auth_clients
SET
    client_id=:client_id,
    client_secret=:client_secret,
    extra=:extra,
    redirect_uri=:redirect_uri,
    updated_at=:updated_at
WHERE id=:id
