SELECT
    a.id,
    a.name,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topic_types a
WHERE a.id=:id AND a.deleted=0