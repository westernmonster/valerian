DELETE FROM oauth_tokens
WHERE expired_at > :expired_at