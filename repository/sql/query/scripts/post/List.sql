SELECT
	p.id,
    p.title,
    p.content,
    p.newsletter_id,
	p.created_at,
    p.updated_at
FROM
	posts as p
ORDER BY p.created_at