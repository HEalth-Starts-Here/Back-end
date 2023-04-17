package noterepository

const (
	// queryCheckUserRole = `
	// SELECT true
	// FROM medics
	// WHERE medics.vkid = $1
	// UNION ALL
	// SELECT false
	// FROM patients
	// WHERE patients.vkid = $1;
	// `

	// queryCreateComment = `
	// INSERT INTO
    // comments (authorismedic, creatingDate, diaryId, text)
	// VALUES
    // (
	// 	$1,
    //     $2,
    //     $3,
    //     $4
    // )
	// RETURNING id, authorismedic, isreaded, creatingDate, text, diaryId;
	// `

	queryGetNote = `
	SELECT $1, id, creatingdate, text
	FROM notes
	WHERE $2 = $3;
	`
	queryGetNote2 = `
	SELECT medicrecordid, id, creatingdate, text
	FROM notes
	WHERE medicrecordid = $1;
	`
	queryGetNote3 = `
	SELECT $1, id, creatingdate, text
	FROM notes
	WHERE medicrecordid = 1;
	`
	// queryDeleteComment = `
	// DELETE FROM comments
	// WHERE id = $1;
	// `
)
