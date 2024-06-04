select auth.users.email users from auth.users
join public.subscriptions subs on subs.user_id = users.id
where subs.newsletter_id = @newsletterId