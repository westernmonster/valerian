SELECT
    a.id,
    a.en_name,
    a.cn_name,
    a.code,
    a.prefix,
    a.deleted,
    a.created_at,
    a.updated_at
FROM country_codes a
WHERE a.id=:id AND a.deleted=0