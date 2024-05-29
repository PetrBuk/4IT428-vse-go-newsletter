INSERT INTO newsletters (created_at, updated_at, name, description, owner_id)
VALUES (NOW(), NOW(), @name, @description, @owner_id)
RETURNING id, name, description, owner_id, created_at, updated_at;
