SELECT
    a.id,
    a.code_type,
    a.used,
    a.code,
    a.identity,
    a.deleted,
    a.created_at,
    a.updated_at
FROM valcodes a
WHERE a.id=:id AND a.deleted=0