SELECT
	n.id,
	n.created_at,
	n.updated_at,
    n.title,
    n.content
FROM
	newsletters as n
WHERE
	id = @id