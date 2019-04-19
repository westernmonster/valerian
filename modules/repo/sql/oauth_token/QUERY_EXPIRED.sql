SELECT COUNT(1) as count
FROM oauth_tokens a
WHERE a.deleted=0 AND expired_at > :expired_at