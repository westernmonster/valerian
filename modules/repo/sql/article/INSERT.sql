INSERT INTO articles(
    id,
    title,
    cover,
    introduction,
    important,
    created_by,
    deleted,
    created_at,
    updated_at
) VALUES (
    :id,
    :title,
    :cover,
    :introduction,
    :important,
    :created_by,
    :deleted,
    :created_at,
    :updated_at
)