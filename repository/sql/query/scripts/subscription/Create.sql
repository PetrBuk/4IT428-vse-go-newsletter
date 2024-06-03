INSERT INTO subscriptions (created_at, newsletter_id, user_id)
SELECT NOW(), @newsletter_id, @user_id
RETURNING id, created_at, newsletter_id, user_id;