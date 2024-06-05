ALTER TABLE subscriptions
DROP COLUMN IF EXISTS is_confirmed;

ALTER TABLE subscriptions
    ADD COLUMN user_id uuid NOT NULL;

ALTER TABLE subscriptions
    ADD CONSTRAINT unique_user_newsletter UNIQUE (user_id, newsletter_id);

ALTER TABLE subscriptions
    ADD CONSTRAINT fk_subscriptions_user_id
        FOREIGN KEY (user_id)
            REFERENCES auth.users (id) ON DELETE CASCADE;

ALTER TABLE subscriptions
DROP COLUMN IF EXISTS email;

ALTER TABLE subscriptions
DROP CONSTRAINT IF EXISTS unique_email_newsletter;
