DELETE FROM subscriptions
WHERE email = @email
AND newsletter_id = @newsletter_id;