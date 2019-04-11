SELECT
    a.id,
    a.topic_id,
    a.account_id,
    a.role,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topic_members a
WHERE a.id=:id AND a.deleted=0