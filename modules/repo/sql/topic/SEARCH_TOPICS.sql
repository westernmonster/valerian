SELECT a.id as topic_id, a.name, a.version_name, a.version_lang
FROM topics a
WHERE a.deleted=0 %s
