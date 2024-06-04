ALTER TABLE subscriptions
    ADD COLUMN is_confirmed BOOLEAN NOT NULL DEFAULT FALSE;

ALTER TABLE subscriptions
DROP CONSTRAINT fk_subscriptions_user_id;

ALTER TABLE subscriptions
DROP COLUMN IF EXISTS user_id;

--ALTER TABLE subscriptions
--DROP CONSTRAINT unique_user_newsletter;

ALTER TABLE subscriptions
ADD COLUMN email text NOT NULL;

ALTER TABLE subscriptions
ADD CONSTRAINT unique_email_newsletter UNIQUE (email, newsletter_id);
