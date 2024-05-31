SELECT
	p.id,
	p.created_at,
	p.title,
	p.content,
    p.newsletter_id,
FROM
	posts as p
WHERE
	id = @id