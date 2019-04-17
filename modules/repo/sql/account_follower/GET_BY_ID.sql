SELECT
    a.id,
    a.account_id,
    a.followers_id,
    a.deleted,
    a.created_at,
    a.updated_at
FROM account_followers a
WHERE a.id=:id AND a.deleted=0