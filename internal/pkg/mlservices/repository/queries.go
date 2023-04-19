package mlservicesrepository

const (
	queryGetAudioList = `
	SELECT filename
	FROM audio;
	`

	queryCreateMedicRecordDiarisation = `
	INSERT INTO 
	audio (medicRecordId, creatingdate, filename)
	VALUES (
		$1,
		$2,
		$3
	)
	RETURNING id, creatingdate, medicRecordId, filename;
	`

	querySetMedicRecordDiarisation = `
	UPDATE audio
	SET diarisation = $2, iscomplete = true
	WHERE id = $1
	RETURNING NULL
	`
)
