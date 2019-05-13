UPDATE topics
SET
    name=:name,
    cover=:cover,
    introduction=:introduction,
    is_private=:is_private,
    allow_chat=:allow_chat,
    allow_discuss=:allow_discuss,
    edit_permission=:edit_permission,
    view_permission=:view_permission,
    join_permission=:join_permission,
    important=:important,
    mute_notification=:mute_notification,
    category_view_type=:category_view_type,
    topic_home=:topic_home,
    topic_type=:topic_type,
    version_name=:version_name,
    created_by=:created_by,
    updated_at=:updated_at
WHERE id=:id
