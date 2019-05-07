UPDATE topic_relations
SET
    from_topic_id=:from_topic_id,
    to_topic_id=:to_topic_id,
    relation=:relation,
    updated_at=:updated_at
WHERE id=:id
