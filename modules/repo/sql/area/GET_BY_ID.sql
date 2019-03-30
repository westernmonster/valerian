SELECT
    a.id,
    a.name,
    a.code,
    a.type,
    a.parent,
    a.deleted,
    a.created_at,
    a.updated_at
FROM areas a
WHERE a.id=:id AND a.deleted=0