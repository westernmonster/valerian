SELECT COUNT(1) as count
FROM areas a
WHERE a.deleted=0 %s