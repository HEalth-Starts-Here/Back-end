package recordrepository

const (
	queryDiaryExist = `
	SELECT 
	FROM diaries
	WHERE id = $1;
	`

	queryMedicRecordExist = `
	SELECT 
	FROM medicrecords
	WHERE id = $1;
	`

	queryMedicExist = `
	SELECT 
	FROM medics
	WHERE vkid = $1;
	`

	//TODO rename to filename
	queryGetImageList = `
	SELECT name
	FROM medicimages
	UNION ALL
	SELECT name
	FROM patientimages;
	`

	queryCreateMedicRecord = `
	INSERT INTO 
	medicrecords (diaryid, creatingdate, title, treatment, recommendations, details)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
	RETURNING id, diaryid, creatingdate, title, treatment, recommendations, details;
	`

	queryCreateMedicRecordImageListFirstPart = `
	INSERT INTO
    medicimages (recordId, name)
	VALUES
	`
    // images (ismedic, recordId, name)

	queryCreatePatientRecordImageListFirstPart = `
	INSERT INTO
    patientimages (recordId, name)
	VALUES
	`
    // patientimages (ismedic, recordId, name)

	queryCreateRecordImageListSecondPart = `
	RETURNING id;
	`

	queryCreateImageTagListFirstPart = `
	INSERT INTO
    tags (imageid, name)
	VALUES
	`

	queryCreateImageTagListSecondPart = `
	RETURNING imageId, name;
	`

	queryGetMedicRecordImageList = `
	SELECT name
	FROM medicimages
	WHERE recordid = $1;
	`
	// WHERE ismedic = $1 AND recordid = $2;

	queryGetPatientRecordImageList = `
	SELECT name
	FROM medicimages
	WHERE recordid = $1;
	`
	// WHERE ismedic = $1 AND recordid = $2;

	queryGetMedicRecordInfo = `
	SELECT diaryid, id, creatingdate, title, treatment, recommendations, details
	FROM medicrecords
	WHERE id = $1;
	`

	queryGetMedicRecordDiarisationList = `
	SELECT id, medicrecordid, creatingdate, diarisation, filename 
	FROM audio
	WHERE medicrecordid = $1;
	`

	queryUpdateTextMedicRecord = `
	UPDATE medicrecords
	SET title = $1, treatment = $2, recommendations = $3, details = $4
	WHERE id = $5
	RETURNING id, diaryid, creatingdate, title, treatment, recommendations, details;
	`

	queryDeleteMedicRecord = `
	DELETE 
	FROM medicrecords
	WHERE id = $1
	RETURNING id;
	`

	queryDeleteImageMedicRecord = `
	DELETE FROM medicimages
	WHERE ismedic = $1 AND recordid = $2
	RETURNING name;
	`

	queryGetBasicUpdateImageMedicRecord = `
	SELECT diaryid, id, creatingdate
	FROM medicrecords
	WHERE id = $1;
	`

	queryGetMedicIdFromDiary = `
	SELECT medicid
	FROM diaries
	WHERE id = $1;
	`

	queryGetMedicIdFromDiaryOfRecord = `
	SELECT medicid
	FROM diaries
	JOIN medicrecords
	ON diaries.id = medicrecords.diaryid
	WHERE medicrecords.id = $1;
	`
	
	queryCreatePatientRecord = `
	INSERT INTO 
	patientrecords (diaryid, creatingdate, title, complaints, treatment, details)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)
	RETURNING id, diaryid, creatingdate, title, complaints, treatment, details;
	`
	
	queryGetPatientRecordInfo = `
	SELECT diaryid, id, creatingdate, title, complaints, treatment, details
	FROM patientrecords
	WHERE id = $1;
	`

	queryUpdateTextPatientRecord = `
	UPDATE patientrecords
	SET title = $1, complaints = $2, treatment = $3, details = $4
	WHERE id = $5
	RETURNING id, diaryid, creatingdate, title, complaints, treatment, details;
	`
)
