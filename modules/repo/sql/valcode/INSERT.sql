INSERT INTO valcodes(
    id,
    code_type,
    used,
    code,
    identity,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :code_type,
    :used,
    :code,
    :identity,
    :deleted,
    :created_at,
    :updated_at
)