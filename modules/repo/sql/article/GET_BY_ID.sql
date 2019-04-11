SELECT
    a.id,
    a.title,
    a.cover,
    a.introduction,
    a.important,
    a.created_by,
    a.deleted,
    a.created_at,
    a.updated_at
FROM articles a
WHERE a.id=:id AND a.deleted=0