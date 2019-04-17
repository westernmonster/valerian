SELECT
    a.id,
    a.locale,
    a.name,
    a.deleted,
    a.created_at,
    a.updated_at
FROM locales a
WHERE a.id=:id AND a.deleted=0