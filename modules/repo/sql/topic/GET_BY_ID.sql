SELECT
    a.id,
    a.name,
    a.cover,
    a.introduction,
    a.is_private,
    a.allow_discuss,
    a.edit_permission,
    a.view_permission,
    a.join_permission,
    a.important,
    a.mute_notification,
    a.category_view_type,
    a.created_by,
    a.deleted,
    a.created_at,
    a.updated_at
FROM topics a
WHERE a.id=:id AND a.deleted=0