SELECT
	n.id,
	n.created_at,
	n.updated_at
FROM
	newsletters as n
WHERE
	id = @id