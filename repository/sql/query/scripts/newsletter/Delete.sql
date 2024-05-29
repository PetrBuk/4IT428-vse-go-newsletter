DELETE from newsletters
WHERE
 id = @id
 and owner_id = @owner_id;