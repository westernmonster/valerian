INSERT INTO topic_categories(
    id,
    topic_id,
    name,
    parent_id,
    created_by,
    seq,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :topic_id,
    :name,
    :parent_id,
    :created_by,
    :seq,
    :deleted,
    :created_at,
    :updated_at
)