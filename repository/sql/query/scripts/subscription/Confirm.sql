update subscriptions
set is_confirmed = TRUE
where newsletter_id =  @newsletter_id
and email = @email
RETURNING id, created_at, newsletter_id, email, is_confirmed;