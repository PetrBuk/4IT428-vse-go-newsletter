INSERT INTO posts (created_at, title, content, newsletter_id)
VALUES (NOW(), @title, @content, @newsletter_id)
RETURNING id, created_at, title, content, newsletter_id;
