ALTER TABLE posts
    ADD COLUMN updated_at     TIMESTAMPTZ   NOT NULL DEFAULT NOW();

-- Drop policies set up for the profiles table
DROP POLICY "Users can update own profile." ON profiles;
DROP POLICY "Users can insert their own profile." ON profiles;
DROP POLICY "Users can view their own profile." ON profiles;

-- Disable Row Level Security on the profiles table
ALTER TABLE profiles DISABLE ROW LEVEL SECURITY;

-- Drop the profiles table
DROP TABLE profiles;