UPDATE accounts
SET
    mobile=:mobile,
    email=:email,
    password=:password,
    role=:role,
    salt=:salt,
    gender=:gender,
    birth_year=:birth_year,
    birth_month=:birth_month,
    birth_day=:birth_day,
    location=:location,
    introduction=:introduction,
    avatar=:avatar,
    source=:source,
    ip=:ip,
    id_cert=:id_cert,
    work_cert=:work_cert,
    is_org=:is_org,
    is_vip=:is_vip,
    updated_at=:updated_at
WHERE id=:id
