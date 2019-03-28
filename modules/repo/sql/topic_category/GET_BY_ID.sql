SELECT
    a.id,
    a.topic_id,
    a.name,
    a.parent_id,
    a.created_by,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topic_categories a
WHERE a.id=:id AND a.deleted=0