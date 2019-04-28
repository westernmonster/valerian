SELECT
    a.id,
    a.name,
    a.emoji,
    a.cn_name,
    a.code,
    a.prefix,
    a.deleted,
    a.created_at,
    a.updated_at
FROM country_codes a
WHERE a.id=:id AND a.deleted=0