DELETE
FROM posts
WHERE id = @id
  AND newsletter_id = @newsletter_id
    AND newsletter_id =
        (SELECT @newsletter_id
        FROM newsletters
        WHERE id = @newsletter_id
        AND owner_id = @user_id)

