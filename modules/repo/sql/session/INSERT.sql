INSERT INTO session(
    id,
    session_type,
    used,
    account_id,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :session_type,
    :used,
    :account_id,
    :deleted,
    :created_at,
    :updated_at
)