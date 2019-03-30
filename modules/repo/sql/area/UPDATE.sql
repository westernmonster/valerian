UPDATE areas
SET
    name=:name,
    code=:code,
    type=:type,
    parent=:parent,
    updated_at=:updated_at
WHERE id=:id
