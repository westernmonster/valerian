SELECT a.*
FROM id_certifications a
WHERE a.deleted=0 %s
ORDER BY a.id DESC