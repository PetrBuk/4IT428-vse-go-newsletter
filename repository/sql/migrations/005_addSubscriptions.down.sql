ALTER TABLE subscriptions
DROP CONSTRAINT fk_subscriptions_user_id;

ALTER TABLE subscriptions
DROP CONSTRAINT fk_subscriptions_newsletter_id;

ALTER TABLE subscriptions
    DROP CONSTRAINT unique_user_newsletter;

DROP TABLE IF EXISTS subscriptions;