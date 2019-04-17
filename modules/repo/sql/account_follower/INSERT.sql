INSERT INTO account_followers(
    id,
    account_id,
    followers_id,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :account_id,
    :followers_id,
    :deleted,
    :created_at,
    :updated_at
)