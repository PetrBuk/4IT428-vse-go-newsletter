SELECT
	p.id,
	n.created_at,
	n.title,
	n.content,
	n.newsletter_id,
FROM
	posts as p
ORDER BY p.created_at