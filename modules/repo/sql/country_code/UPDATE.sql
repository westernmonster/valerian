UPDATE country_codes
SET
    name=:name,
    cn_name=:cn_name,
    code=:code,
    emoji=:emoji,
    prefix=:prefix,
    updated_at=:updated_at
WHERE id=:id
