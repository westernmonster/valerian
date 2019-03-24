UPDATE valcodes
SET
    code_type=:code_type,
    used=:used,
    code=:code,
    identity=:identity,
    updated_at=:updated_at
WHERE id=:id
