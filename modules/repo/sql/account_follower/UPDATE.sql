UPDATE account_followers
SET
    account_id=:account_id,
    followers_id=:followers_id,
    updated_at=:updated_at
WHERE id=:id
