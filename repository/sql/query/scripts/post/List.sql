SELECT
	p.id,
    p.title,
    p.content,
    p.newsletter_id,
	p.created_at,
    p.updated_at,
    p.is_published
FROM
	posts as p
ORDER BY p.created_at