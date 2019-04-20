SELECT
a.*
FROM auth_clients a
WHERE a.deleted=0 %s