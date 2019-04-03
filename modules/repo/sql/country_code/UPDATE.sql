UPDATE country_codes
SET
    en_name=:en_name,
    cn_name=:cn_name,
    code=:code,
    prefix=:prefix,
    updated_at=:updated_at
WHERE id=:id
