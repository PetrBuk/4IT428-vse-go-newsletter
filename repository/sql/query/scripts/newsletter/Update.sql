UPDATE newsletters
SET
    title = @title,
    content = @content,
    updated_at = now()
WHERE
   id = @id