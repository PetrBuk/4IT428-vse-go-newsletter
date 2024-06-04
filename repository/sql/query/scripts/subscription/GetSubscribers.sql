select subs.email
form public.subscriptions subs
where subs.newsletter_id = @newsletterId
and subs.is_confirmed = true;