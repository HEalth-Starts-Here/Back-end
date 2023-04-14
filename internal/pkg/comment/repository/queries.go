package commentrepository

const (
	queryCheckUserRole = `
	SELECT true
	FROM medics
	WHERE medics.vkid = $1
	UNION ALL
	SELECT false
	FROM patients
	WHERE patients.vkid = $1;
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
	RETURNING id, authorismedic, creatingDate, diaryId, text;
	`
	
	// queryLinkDiary = `
	// UPDATE diaries
	// SET patientid = $2
	// FROM diaries d
	// JOIN medics m
	// ON d.medicid = m.vkid
	// WHERE diaries.id = $1
	// RETURNING d.id, d.medicid, m.name, d.patientid, d.creatingdate, d.title, d.complaints, d.anamnesis, d.objectively, d.diagnosis;
	// `

	// queryDeleteDiary = `
	// DELETE FROM diaries
	// WHERE id = $1;
	// `

	// queryDiaryList = `
	// SELECT id, medicid, medics.name, patientid, patients.name, creatingdate, title, objectively
	// FROM diaries
	// LEFT JOIN patients ON diaries.patientid = patients.vkid
	// JOIN medics ON diaries.medicid = medics.vkid
	// WHERE medicid = $1 OR patientid = $1
	// ORDER BY creatingdate;
	// `

	// queryGetCertainDiaryMainInfo = `
	// SELECT patients.name, diaries.id, medicid, medics.name, patientid, creatingDate, title, complaints, anamnesis, objectively, diagnosis
	// FROM diaries 
	// LEFT JOIN patients on diaries.patientid = patients.vkid
	// LEFT JOIN medics on diaries.medicid = medics.vkid
	// WHERE diaries.id = $1;
	// `

	// // queryGetCertainDiaryRecords = `
	// // SELECT (id, diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness)
	// // FROM records
	// // WHERE diaryid = $1;
	// // `
	// queryGetCertainDiaryMedicRecords = `
	// SELECT id, creatingdate, title, details
	// FROM medicRecords
	// WHERE diaryid = $1;
	// `

	// queryGetCertainDiaryPatientRecords = `
	// SELECT id, creatingdate, title, details
	// FROM patientRecords
	// WHERE diaryid = $1;
	// `

	// queryGetUserRole = `
	// SELECT true
	// FROM medics 
	// WHERE vkid = $1 
	// UNION ALL 
	// SELECT false
	// FROM patients 
	// WHERE vkid = $1;
	// `

	// queryGetCertainDiaryRecords = `
	// SELECT diaries.medicid as userid, mr.creatingdate, mr.title, details
	// FROM medicRecords mr
	// JOIN diaries ON diaries.id = diaryid 
	// WHERE diaryid = $1
	// UNION ALL
	// SELECT diaries.patientid as userid, pr.creatingdate, pr.title, details
	// FROM patientRecords pr
	// JOIN diaries ON diaries.id = diaryid 
	// WHERE diaryid = $1;
	// `

	// // queryCreateRecord = `
	// // INSERT INTO
    // // records (diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness)
	// // VALUES
    // // (
    // //     $1,
    // //     $2,
    // //     $3,
    // //     $4,
    // //     $5,
    // //     $6,
    // //     $7,
    // //     $8,
    // //     $9,
    // //     $10,
    // //     $11
    // // )
	// // RETURNING id, diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness;
	// // `

	// queryUpdateDiary = `
	// UPDATE diaries
	// SET title = $1, complaints = $2, anamnesis = $3, objectively = $4, diagnosis = $5
	// WHERE id = $6
	// RETURNING id, medicid, patientid, creatingdate, title, complaints, anamnesis, objectively, diagnosis;
	// `
)
