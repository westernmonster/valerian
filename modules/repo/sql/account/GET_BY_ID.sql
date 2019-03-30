SELECT
    a.id,
    a.mobile,
    a.email,
    a.password,
    a.gender,
    a.birth_year,
    a.birth_month,
    a.birth_day,
    a.introduction,
    a.location,
    a.avatar,
    a.source,
    a.ip,
    a.deleted,
    a.created_at,
    a.updated_at
FROM accounts a
WHERE a.id=:id AND a.deleted=0