SELECT
a.*
FROM auth_access a
WHERE a.deleted=0 %s