package diaryrepository

// import "fmt"

// func encryptVKID(id uint64)(string){
// 	return fmt.Sprintf("encrypt(%d::text::bytea,'secret','aes')", id)
// }

const (
	// encryptVKID = "encrypt(%s::text::bytea,'secret','aes')"

	queryCreateDiary = `
	INSERT INTO
    diaries (medicId, creatingDate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate)
	VALUES
    (
		encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text,
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
	RETURNING id, convert_from(decrypt(medicId::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, creatingDate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate;
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
	SET patientid = encrypt($2::bigint::text::bytea,'secret'::bytea,'aes'::text)::text
	FROM diaries d
	JOIN medics m
	ON d.medicid = m.vkid
	WHERE diaries.id = $1
	RETURNING d.id, convert_from(decrypt(d.medicid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, m.name, convert_from(decrypt(d.patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, d.creatingdate, d.title, d.complaints, d.anamnesis, d.objectively, d.diagnosis, d.variant, d.frequency, d.startdate;
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
	SELECT id, convert_from(decrypt(medicid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, medics.name, convert_from(decrypt(patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, patients.name, creatingdate, title, objectively, diarytokens.token, iscomplete
	FROM diaries
	LEFT JOIN patients ON diaries.patientid = patients.vkid
	JOIN medics ON diaries.medicid = medics.vkid
	LEFT JOIN diarytokens ON diaries.id = diarytokens.diaryid
	WHERE medicid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text OR patientid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text
	ORDER BY creatingdate DESC;
	`

	queryGetCertainDiaryMainInfo = `
	SELECT patients.name, diaries.id, convert_from(decrypt(medicid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, medics.name, convert_from(decrypt(patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, creatingDate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate
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
	// queryGetUserRole = fmt.Sprintf(`
	// SELECT true
	// FROM medics
	// WHERE %s = $1
	// UNION ALL
	// SELECT false
	// FROM patients
	// WHERE %s = $1;
	// `, queryGetCertainDiaryPatientRecords, queryGetCertainDiaryPatientRecords)

	queryGetUserRole = `
	SELECT true
	FROM medics 
	WHERE vkid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text 
	UNION ALL 
	SELECT false
	FROM patients 
	WHERE vkid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text;
	`

	queryGetCertainDiaryRecords = `
	SELECT convert_from(decrypt(diaries.medicid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint as userid, mr.creatingdate, mr.title, details
	FROM medicRecords mr
	JOIN diaries ON diaries.id = diaryid 
	WHERE diaryid = $1
	UNION ALL
	SELECT convert_from(decrypt(diaries.patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint as userid, pr.creatingdate, pr.title, details
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
	RETURNING id, convert_from(decrypt(medicid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, convert_from(decrypt(patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, creatingdate, title, complaints, anamnesis, objectively, diagnosis, variant, frequency, startdate;
	`
)
