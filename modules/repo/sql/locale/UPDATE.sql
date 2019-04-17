UPDATE locales
SET
    locale=:locale,
    name=:name,
    updated_at=:updated_at
WHERE id=:id
