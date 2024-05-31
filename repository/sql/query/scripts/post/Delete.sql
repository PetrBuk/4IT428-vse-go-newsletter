DELETE from posts
WHERE
 id = @id
 and newsletter_id = @newsletter_id;