INSERT INTO subscriptions (created_at, newsletter_id, email, is_confirmed)
values(NOW(), @newsletter_id, @email, FALSE)
RETURNING id, created_at, newsletter_id, email, is_confirmed;