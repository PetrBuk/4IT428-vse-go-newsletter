DELETE FROM posts
WHERE id = @id
  AND newsletter_id =
      (SELECT id
       FROM newsletters
       WHERE id = posts.newsletter_id
         AND owner_id = @user_id);
