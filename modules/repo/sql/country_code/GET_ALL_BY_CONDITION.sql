SELECT a.*
FROM country_codes a
WHERE a.deleted=0 %s
ORDER BY a.id DESC