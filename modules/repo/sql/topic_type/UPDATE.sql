UPDATE topic_types
SET
    name=:name,
    updated_at=:updated_at
WHERE id=:id
