INSERT INTO accounts(
    id,
    mobile,
    email,
    password,
    role,
    salt,
    gender,
    birth_year,
    birth_month,
    birth_day,
    location,
    introduction,
    avatar,
    source,
    ip,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :mobile,
    :email,
    :password,
    :role,
    :salt,
    :gender,
    :birth_year,
    :birth_month,
    :birth_day,
    :location,
    :introduction,
    :avatar,
    :source,
    :ip,
    :deleted,
    :created_at,
    :updated_at
)