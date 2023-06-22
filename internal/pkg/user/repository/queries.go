package userrepository

const (
	queryGetUserInfo = `
	SELECT vkid, name, true 
	FROM medics 
	WHERE vkid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text 
	UNION ALL 
	SELECT vkid, name, false 
	FROM patients 
	WHERE vkid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text;
	`

	queryRegisterMedic = `
	INSERT INTO
    medics (vkid, name)
	VALUES
    (
		encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text,
        $2
    )
	RETURNING convert_from(decrypt(vkid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, name;
	`

	queryRegisterPatient = `
	INSERT INTO
    patients (vkid, name)
	VALUES
    (
		encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text,
        $2
    )
	RETURNING convert_from(decrypt(vkid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, name;
	`

	queryLinkPatientToDiary = `
	UPDATE diaries
	SET patientid = encrypt($1::bigint::text::bytea,'secret'::bytea,'aes'::text)::text
	WHERE id = $2
	RETURNING convert_from(decrypt(patientid::text::bytea,'secret','aes'),'SQL_ASCII')::bigint, id;
	`

	queryDeleteLinkToken = `
	DELETE FROM diarytokens
	WHERE diaryid = $1 AND token = $2
	RETURNING true;
	`

// 	INSERT INTO
//     medics (vkid, name)
// VALUES
//     (
//         111,
//         'Иван Иванов'
//     ),
//     (
//         222,
//         'Петр Петров'
//     );

// queryMedicInit = `
// INSERT INTO
// medics (vkid, name)
// VALUES
// (
// 	$1,
//     $2
// )
// RETURNING vkid, name;
// `

// queryPatientInit = `
// INSERT INTO
// patients (vkid, name)
// VALUES
// (
// 	$1,
//     $2
// )
// RETURNING vkid, name;
// `
)
