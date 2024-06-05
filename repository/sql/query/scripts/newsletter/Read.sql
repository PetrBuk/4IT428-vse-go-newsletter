SELECT
	n.id,
	n.created_at,
	n.updated_at,
    n.name,
    n.description,
    n.owner_id
FROM
	newsletters as n
WHERE
	id = @id