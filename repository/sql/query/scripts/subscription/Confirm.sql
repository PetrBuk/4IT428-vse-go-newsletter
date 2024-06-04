update subscriptions
set is_confirmed = TRUE
where newsletter_id =  @newsletter_id
and user_id = @user_id
RETURNING id, created_at, newsletter_id, user_id, is_confirmed;