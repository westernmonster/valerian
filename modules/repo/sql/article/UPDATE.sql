UPDATE articles
SET
    title=:title,
    cover=:cover,
    introduction=:introduction,
    important=:important,
    created_by=:created_by,
    updated_at=:updated_at
WHERE id=:id
