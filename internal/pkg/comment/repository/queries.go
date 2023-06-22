package commentrepository

const (
	queryCheckUserRole = `
	SELECT true
	FROM medics
	WHERE medics.vkid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text
	UNION ALL
	SELECT false
	FROM patients
	WHERE patients.vkid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text;
	`

	queryCreateComment = `
	INSERT INTO
    comments (authorismedic, creatingDate, diaryId, text)
	VALUES
    (
		$1,
        $2,
        $3,
        $4
    )
	RETURNING id, authorismedic, isreaded, creatingDate, text, diaryId;
	`

	queryGetComment = `
	SELECT id, authorismedic, isreaded, creatingdate, text, diaryid
	FROM comments
	WHERE diaryid = $1;
	`
	queryDeleteComment = `
	DELETE FROM comments
	WHERE id = $1;
	`
)

