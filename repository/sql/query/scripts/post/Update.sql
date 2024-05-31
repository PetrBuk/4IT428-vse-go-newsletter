UPDATE posts
SET
    title = @title,
    content = @content,
    updated_at = now()
WHERE
   id = @id
   and newsletter_id = @newsletter_id
RETURNING id, title, content, newsletter_id;