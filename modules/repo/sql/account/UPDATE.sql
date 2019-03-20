UPDATE accounts
SET
    mobile=:mobile,
    email=:email,
    password=:password,
    gender=:gender,
    birth_year=:birth_year,
    birth_month=:birth_month,
    birth_day=:birth_day,
    introduction=:introduction,
    avatar=:avatar,
    source=:source,
    ip=:ip,
    updated_at=:updated_at
WHERE id=:id
