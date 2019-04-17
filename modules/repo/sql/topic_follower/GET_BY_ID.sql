SELECT
    a.id,
    a.topic_id,
    a.followers_id,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topic_followers a
WHERE a.id=:id AND a.deleted=0