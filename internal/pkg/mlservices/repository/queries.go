package mlservicesrepository

const (
	queryGetAudioList = `
	SELECT filename
	FROM audio;
	`

	queryCreateMedicRecordDiarisation = `
	INSERT INTO 
	audio (medicRecordId, creatingdate, diarisation, filename)
	VALUES (
		$1,
		$2,
		$3,
		$4
	)
	RETURNING id, creatingdate, medicRecordId, diarisation, filename;
	`
)
