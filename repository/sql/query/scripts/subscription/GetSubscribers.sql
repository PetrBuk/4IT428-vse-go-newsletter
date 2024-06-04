SELECT email
FROM subscriptions
WHERE newsletter_id = @newsletterId
  AND is_confirmed = true;
