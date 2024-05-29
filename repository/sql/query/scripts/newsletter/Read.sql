SELECT
	n.id,
	n.created_at,
	n.updated_at,
    n.name,
    n.description
FROM
	newsletters as n
WHERE
	id = @id