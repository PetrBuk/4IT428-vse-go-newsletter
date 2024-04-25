DROP TABLE IF EXISTS newsletters;

-- Drop the trigger that handles new user profiles creation
DROP TRIGGER on_auth_user_created ON auth.users;

-- Drop the function used by the trigger
DROP FUNCTION public.handle_new_user();

-- Drop policies set up for the profiles table
DROP POLICY "Users can update own profile." ON profiles;
DROP POLICY "Users can insert their own profile." ON profiles;
DROP POLICY "Users can view their own profile." ON profiles;

-- Disable Row Level Security on the profiles table
ALTER TABLE profiles DISABLE ROW LEVEL SECURITY;

-- Drop the profiles table
DROP TABLE profiles;