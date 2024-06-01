INSERT INTO posts (created_at, updated_at, title, content, newsletter_id)
SELECT NOW(), NOW(), @title, @content, @newsletter_id
FROM newsletters
WHERE id = @newsletter_id
  AND owner_id = @user_id
RETURNING id, created_at, updated_at, title, content, newsletter_id;