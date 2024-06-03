CREATE TABLE subscriptions (
                               id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                               created_at TIMESTAMPTZ NOT NULL,
                               user_id uuid NOT NULL,
                               newsletter_id uuid NOT NULL,
                               CONSTRAINT fk_subscriptions_user_id
                                   FOREIGN KEY (user_id)
                                       REFERENCES auth.users (id) ON DELETE CASCADE,
                               CONSTRAINT fk_subscriptions_newsletter_id
                                   FOREIGN KEY (newsletter_id)
                                       REFERENCES newsletters (id) ON DELETE CASCADE,
                               CONSTRAINT unique_user_newsletter UNIQUE (user_id, newsletter_id)
);
