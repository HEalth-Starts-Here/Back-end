package searchrepository

const (
	// queryUserExist = `
	// SELECT 
	// FROM medics
	// WHERE vkid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text
	// UNION ALL
	// SELECT
	// FROM patients
	// WHERE vkid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text;
	// `
	
	queryCheckUserRole = `
	SELECT true
	FROM medics
	WHERE medics.vkid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text
	UNION ALL
	SELECT false
	FROM patients
	WHERE patients.vkid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text;
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
	SELECT id, convert_from(decrypt(medicid::text::bytea,'%s','aes'),'SQL_ASCII')::bigint, medics.name, convert_from(decrypt(patientid::text::bytea,'%s','aes'),'SQL_ASCII')::bigint, patients.name, creatingdate, title, objectively, diarytokens.token, iscomplete
	FROM diaries
	LEFT JOIN patients ON diaries.patientid = patients.vkid
	JOIN medics ON diaries.medicid = medics.vkid
	LEFT JOIN diarytokens ON diaries.id = diarytokens.diaryid
	WHERE (medicid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text OR patientid = encrypt($1::bigint::text::bytea,'%s'::bytea,'aes'::text)::text) AND 
	(LOWER(title) LIKE LOWER('%s') OR LOWER(objectively) LIKE LOWER('%s') OR LOWER(%s.name) LIKE LOWER('%s'))
	ORDER BY creatingdate DESC;
	`
)
