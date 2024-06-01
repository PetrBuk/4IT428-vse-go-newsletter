UPDATE posts
SET
    title = @title,
    content = @content,
    updated_at = NOW()
    FROM newsletters
WHERE
    posts.newsletter_id = newsletters.id
  AND newsletters.id = @newsletter_id
  AND newsletters.owner_id = @user_id
  AND posts.id = @id
  AND posts.newsletter_id = @newsletter_id
  AND posts.is_published = FALSE
    RETURNING posts.id, posts.title, posts.content, posts.newsletter_id, posts.created_at, posts.updated_at, posts.is_published;
