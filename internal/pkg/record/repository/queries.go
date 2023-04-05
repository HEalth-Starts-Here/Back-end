package recordrepository

const (
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


)
