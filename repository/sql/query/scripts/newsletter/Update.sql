UPDATE newsletters
SET
    name = @name,
    description = @description,
    updated_at = now()
WHERE
   id = @id
   and owner_id = @owner_id
RETURNING id, name, description, owner_id, updated_at;