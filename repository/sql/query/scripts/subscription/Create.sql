INSERT INTO subscriptions (created_at, newsletter_id, user_id, is_confirmed)
values(NOW(), @newsletter_id, @user_id, FALSE)
RETURNING id, created_at, newsletter_id, user_id, is_confirmed;