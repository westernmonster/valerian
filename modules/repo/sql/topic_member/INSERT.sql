INSERT INTO topic_members(
    id,
    topic_id,
    account_id,
    role,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :topic_id,
    :account_id,
    :role,
    :deleted,
    :created_at,
    :updated_at
)