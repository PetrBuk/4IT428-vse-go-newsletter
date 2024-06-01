SELECT
	p.id,
	p.created_at,
	p.updated_at,
	p.title,
	p.content,
    p.newsletter_id
FROM
	posts as p
WHERE
	id = @id