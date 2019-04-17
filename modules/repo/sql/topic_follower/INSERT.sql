INSERT INTO topic_followers(
    id,
    topic_id,
    followers_id,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :topic_id,
    :followers_id,
    :deleted,
    :created_at,
    :updated_at
)