SELECT
	p.id,
	p.created_at,
	p.updated_at,
	p.title,
	p.content,
    p.newsletter_id,
    p.is_published
FROM
	posts as p
WHERE
	id = @id