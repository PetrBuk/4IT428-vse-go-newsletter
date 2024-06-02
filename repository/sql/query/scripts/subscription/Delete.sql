DELETE FROM subscriptions
WHERE user_id = @user_id
AND newsletter_id = @newsletter_id;