INSERT INTO accounts(
    id,
    mobile,
    email,
    password,
    gender,
    birth_year,
    birth_month,
    birth_day,
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
    :gender,
    :birth_year,
    :birth_month,
    :birth_day,
    :introduction,
    :avatar,
    :source,
    :ip,
    :deleted,
    :created_at,
    :updated_at
)