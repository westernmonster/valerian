SELECT
    a.id,
    a.from_topic_id,
    a.to_topic_id,
    a.relation,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topic_relations a
WHERE a.id=:id AND a.deleted=0