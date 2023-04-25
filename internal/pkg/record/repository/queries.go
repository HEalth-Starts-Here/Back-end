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

	queryUserExist = `
	SELECT 
	FROM medics
	WHERE vkid = $1
	UNION ALL
	SELECT
	FROM patients
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
	FROM patientimages
	WHERE recordid = $1;
	`
	// WHERE ismedic = $1 AND recordid = $2;

	queryGetMedicRecordInfo = `
	SELECT diaryid, id, creatingdate, title, treatment, recommendations, details
	FROM medicrecords
	WHERE id = $1;
	`

	queryGetMedicRecordDiarisationList = `
	SELECT id, medicrecordid, creatingdate, diarisation, filename, iscomplete
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

	queryDeletePatientRecord = `
	DELETE 
	FROM patientrecords
	WHERE id = $1
	RETURNING id;
	`

	queryDeleteImageMedicRecord = `
	DELETE FROM medicimages
	WHERE recordid = $1
	RETURNING name;
	`

	queryDeleteImagePatientRecord = `
	DELETE FROM medicimages
	WHERE recordid = $1
	RETURNING name;
	`

	queryGetBasicUpdateImageMedicRecord = `
	SELECT diaryid, id, creatingdate
	FROM medicrecords
	WHERE id = $1;
	`

	queryGetBasicUpdateImagePatientRecord = `
	SELECT diaryid, id, creatingdate
	FROM patientrecords
	WHERE id = $1;
	`

	queryGetMedicIdFromDiary = `
	SELECT medicid
	FROM diaries
	WHERE id = $1;
	`

	queryGetMedicAndPatientIdsFromDiaryOfRecord = `
	SELECT medicid, patientid
	FROM diaries
	JOIN medicrecords
	ON diaries.id = medicrecords.diaryid
	WHERE medicrecords.id = $1;
	`

	queryCreatePatientRecord = `
	INSERT INTO 
	patientrecords (diaryid, creatingdate, title, complaints, treatment, details, feelings)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)
	RETURNING id, diaryid, creatingdate, title, complaints, treatment, details, feelings;
	`

	queryGetPatientRecordInfo = `
	SELECT diaryid, id, creatingdate, title, complaints, treatment, details, feelings
	FROM patientrecords
	WHERE id = $1;
	`

	queryUpdateTextPatientRecord = `
	UPDATE patientrecords
	SET title = $1, complaints = $2, treatment = $3, details = $4, feelings = $5
	WHERE id = $6
	RETURNING id, diaryid, creatingdate, title, complaints, treatment, details, feelings;
	`

)
