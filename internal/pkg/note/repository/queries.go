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

	// queryCreateNote = `
	// INSERT INTO
    // notes (medicrecordid, ismedicrecord, creatingdate, text)
	// VALUES
    // (
    //     $1,
    //     $2,
    //     $3,
    //     $4
    // )
	// RETURNING id, $5, ismedicrecord, creatingDate, text;
	// `

	queryCreateNoteFirstPart = `
	INSERT INTO
    notes (
	`

	queryCreateNoteSecondPart = `, ismedicrecord, creatingdate, text)
	VALUES
    (
        $1,
        $2,
        $3,
        $4
    )
	RETURNING id, 
	`

	queryCreateNoteThirdPart = `, ismedicrecord, creatingDate, text;	`

	// medicrecordid = "medicrecordid"
	// patientrecordid = "patientrecordid"

	queryGetNoteFirstPart = `
	SELECT 
	`

	queryGetNoteSecondPart = `
	, id, creatingdate, text
	FROM notes
	WHERE `

	queryGetNoteThirdPart = `
	= $1;
	`

	queryDeleteNote = `
	DELETE FROM notes
	WHERE id = $1;
	`
)
