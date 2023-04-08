package recordrepository

const (
	queryDiaryExist = `
	SELECT 
	FROM diaries
	WHERE id = $1;
	`

	queryGetImageList = `
	SELECT name
	FROM images;
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

	queryCreateRecordImageListFirstPart = `
	INSERT INTO
    images (ismedic, recordId, name)
	VALUES
	`

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

	queryGetRecordImageList = `
	SELECT name
	FROM images
	WHERE ismedic = $1 AND recordid = $2;
	`

	queryGetMedicRecordInfo = `
	SELECT diaryid, id, creatingdate, title, treatment, recommendations, details
	FROM medicrecords
	WHERE id = $1;
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
	DELETE FROM images
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

)
