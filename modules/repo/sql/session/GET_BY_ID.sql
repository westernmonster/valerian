SELECT
    a.id,
    a.session_type,
    a.used,
    a.account_id,
    a.deleted,
    a.created_at,
    a.updated_at
FROM session a
WHERE a.id=:id AND a.deleted=0