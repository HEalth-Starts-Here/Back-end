package searchrepository

const (
	queryUserExist = `
	SELECT 
	FROM medics
	WHERE vkid = $1
	UNION ALL
	SELECT
	FROM patients
	WHERE vkid = $1;
	`
	
	queryCheckUserRole = `
	SELECT true
	FROM medics
	WHERE medics.vkid = $1
	UNION ALL
	SELECT false
	FROM patients
	WHERE patients.vkid = $1;
	`

	// querySearchDiary = `
	// SELECT id, medicid, medics.name, patientid, patients.name, creatingdate, title, objectively, diarytokens.token, iscomplete
	// FROM diaries
	// LEFT JOIN patients ON diaries.patientid = patients.vkid
	// JOIN medics ON diaries.medicid = medics.vkid
	// LEFT JOIN diarytokens ON diaries.id = diarytokens.diaryid
	// WHERE (medicid = $1 OR patientid = $1) AND 
	// (title LIKE '%s$2%s' OR objectively LIKE '%s$2%s' OR %s.name LIKE '%s$2%s')
	// ORDER BY creatingdate;
	// `

	querySearchDiary = `
	SELECT id, medicid, medics.name, patientid, patients.name, creatingdate, title, objectively, diarytokens.token, iscomplete
	FROM diaries
	LEFT JOIN patients ON diaries.patientid = patients.vkid
	JOIN medics ON diaries.medicid = medics.vkid
	LEFT JOIN diarytokens ON diaries.id = diarytokens.diaryid
	WHERE (medicid = $1 OR patientid = $1) AND 
	(title LIKE '%s' OR objectively LIKE '%s' OR %s.name LIKE '%s')
	ORDER BY creatingdate;
	`
)
