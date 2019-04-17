SELECT COUNT(1) as count
FROM locales a
WHERE a.deleted=0 %s