package diaryrepository

const (
	queryCreateDiary = `
	INSERT INTO
    diaries (medicId, creatingDate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate)
	VALUES
    (
		$1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9,
        $10
    )
	RETURNING id, medicId, creatingDate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate;
	`

	queryCreateDiaryLinkToken = `
	INSERT INTO
    diarytokens (diaryId, token)
	VALUES
    (
		$1,
        $2
    );
	`

	queryDeleteLinkToken = `
	DELETE FROM diarytokens
	WHERE diaryid = $1 AND token = $2
	RETURNING true;
	`

	queryLinkDiary = `
	UPDATE diaries
	SET patientid = $2
	FROM diaries d
	JOIN medics m
	ON d.medicid = m.vkid
	WHERE diaries.id = $1
	RETURNING d.id, d.medicid, m.name, d.patientid, d.creatingdate, d.title, d.complaints, d.anamnesis, d.objectively, d.diagnosis, d.variant, d.frequency, d.startdate;
	`

	// queryLinkDiary2 = `
	// UPDATE diaries
	// SET patientid = $2
	// FROM diaries d
	// JOIN medics m
	// ON d.medicid = m.vkid
	// WHERE diaries.id = $1;
	// SELECT d.id, d.medicid, m.name, d.patientid, d.creatingdate, d.title, d.complaints, d.anamnesis, d.objectively, d.diagnosis
	// FROM diaries d
	// JOIN medics m
	// ON d.medicid = m.vkid
	// WHERE patientid = $2 AND  d.id = $1;
	// `
	queryDeleteDiary = `
	DELETE FROM diaries
	WHERE id = $1;
	`

	queryCompleteDiary = `
	UPDATE diaries 
	SET iscomplete = true
	WHERE id = $1;
	`

	queryDiaryList = `
	SELECT id, medicid, medics.name, patientid, patients.name, creatingdate, title, objectively, diarytokens.token, iscomplete
	FROM diaries
	LEFT JOIN patients ON diaries.patientid = patients.vkid
	JOIN medics ON diaries.medicid = medics.vkid
	LEFT JOIN diarytokens ON diaries.id = diarytokens.diaryid
	WHERE medicid = $1 OR patientid = $1
	ORDER BY creatingdate DESC;
	`

	queryGetCertainDiaryMainInfo = `
	SELECT patients.name, diaries.id, medicid, medics.name, patientid, creatingDate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate
	FROM diaries 
	LEFT JOIN patients on diaries.patientid = patients.vkid
	LEFT JOIN medics on diaries.medicid = medics.vkid
	WHERE diaries.id = $1;
	`

	// queryGetCertainDiaryRecords = `
	// SELECT (id, diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness)
	// FROM records
	// WHERE diaryid = $1;
	// `
	queryGetCertainDiaryMedicRecords = `
	SELECT id, creatingdate, title, details
	FROM medicRecords
	WHERE diaryid = $1;
	`

	queryGetCertainDiaryPatientRecords = `
	SELECT id, creatingdate, title, details, feelings
	FROM patientRecords
	WHERE diaryid = $1;
	`

	queryGetUserRole = `
	SELECT true
	FROM medics 
	WHERE vkid = $1 
	UNION ALL 
	SELECT false
	FROM patients 
	WHERE vkid = $1;
	`

	queryGetCertainDiaryRecords = `
	SELECT diaries.medicid as userid, mr.creatingdate, mr.title, details
	FROM medicRecords mr
	JOIN diaries ON diaries.id = diaryid 
	WHERE diaryid = $1
	UNION ALL
	SELECT diaries.patientid as userid, pr.creatingdate, pr.title, details
	FROM patientRecords pr
	JOIN diaries ON diaries.id = diaryid 
	WHERE diaryid = $1;
	`

	// queryCreateRecord = `
	// INSERT INTO
	// records (diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness)
	// VALUES
	// (
	//     $1,
	//     $2,
	//     $3,
	//     $4,
	//     $5,
	//     $6,
	//     $7,
	//     $8,
	//     $9,
	//     $10,
	//     $11
	// )
	// RETURNING id, diaryId, creatingDate, title, description, area, dryness, edema, itching, pain, peeling, redness;
	// `

	queryUpdateDiary = `
	UPDATE diaries
	SET title = $1, complaints = $2, anamnesis = $3, objectively = $4, diagnosis = $5, variant = $6, frequency = $7, startdate = $8
	WHERE id = $9
	RETURNING id, medicid, patientid, creatingdate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate;
	`
)
