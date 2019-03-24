UPDATE session
SET
    session_type=:session_type,
    used=:used,
    account_id=:account_id,
    updated_at=:updated_at
WHERE id=:id
