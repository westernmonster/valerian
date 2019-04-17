UPDATE topic_followers
SET
    topic_id=:topic_id,
    followers_id=:followers_id,
    updated_at=:updated_at
WHERE id=:id
