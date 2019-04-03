SELECT COUNT(1) as count
FROM country_codes a
WHERE a.deleted=0 %s