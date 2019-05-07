INSERT INTO topic_relations(
    id,
    from_topic_id,
    to_topic_id,
    relation,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :from_topic_id,
    :to_topic_id,
    :relation,
    :deleted,
    :created_at,
    :updated_at
)