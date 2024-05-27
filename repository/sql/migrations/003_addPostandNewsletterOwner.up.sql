ALTER TABLE newsletters
    ADD COLUMN owner_id uuid;

ALTER TABLE newsletters
ADD CONSTRAINT fk_owner
FOREIGN KEY (owner_id)
REFERENCES profiles (id)
ON DELETE CASCADE;

ALTER TABLE profiles
    ADD COLUMN email text;

CREATE TABLE posts (
id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
created_at     TIMESTAMPTZ   NOT NULL,
title          TEXT          NOT NULL,
content        TEXT          NOT NULL,
newsletter_id  uuid          NOT NULL,
CONSTRAINT fk_newsletter
FOREIGN KEY (newsletter_id)
REFERENCES newsletters (id)
);
