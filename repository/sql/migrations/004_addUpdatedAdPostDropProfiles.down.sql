ALTER TABLE posts
DROP COLUMN IF EXISTS updated_at ;


- Create a table for public profiles
create table profiles (
                          id uuid references auth.users on delete cascade not null primary key,
                          updated_at timestamp with time zone,
                          username text unique,
                          full_name text,
                          avatar_url text,

                          constraint username_length check (char_length(username) >= 3)
);
-- Set up Row Level Security (RLS)
-- See https://supabase.com/docs/guides/auth/row-level-security for more details.
alter table profiles
    enable row level security;

create policy "Users can view their own profile." on profiles
  for select using ((select auth.uid() = id));

create policy "Users can insert their own profile." on profiles
  for insert with check ((select auth.uid()) = id);

create policy "Users can update own profile." on profiles
  for update using ((select auth.uid()) = id);

-- This trigger automatically creates a profile entry when a new user signs up via Supabase Auth.
-- See https://supabase.com/docs/guides/auth/managing-user-data#using-triggers for more details.
create function public.handle_new_user()
    returns trigger as $$
begin
insert into public.profiles (id, full_name, avatar_url)
values (new.id, new.raw_user_meta_data->>'full_name', new.raw_user_meta_data->>'avatar_url');
return new;
end;
$$ language plpgsql security definer;
create trigger on_auth_user_created
    after insert on auth.users
    for each row execute procedure public.handle_new_user();