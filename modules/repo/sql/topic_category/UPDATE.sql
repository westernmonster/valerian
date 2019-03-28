UPDATE topic_categories
SET
    topic_id=:topic_id,
    name=:name,
    parent_id=:parent_id,
    created_by=:created_by,
    updated_at=:updated_at
WHERE id=:id
