INSERT INTO country_codes(
    id,
    en_name,
    cn_name,
    code,
    prefix,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :en_name,
    :cn_name,
    :code,
    :prefix,
    :deleted,
    :created_at,
    :updated_at
)