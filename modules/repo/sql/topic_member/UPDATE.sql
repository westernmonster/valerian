UPDATE topic_members
SET
    topic_id=:topic_id,
    account_id=:account_id,
    role=:role,
    updated_at=:updated_at
WHERE id=:id
