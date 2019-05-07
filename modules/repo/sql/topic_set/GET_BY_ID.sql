SELECT
    a.id,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topic_sets a
WHERE a.id=:id AND a.deleted=0