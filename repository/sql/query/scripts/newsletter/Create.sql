INSERT INTO newsletters (created_at, updated_at, name, description, owner_id)
VALUES (NOW(), NOW(), @name, @description, @owner_id)
