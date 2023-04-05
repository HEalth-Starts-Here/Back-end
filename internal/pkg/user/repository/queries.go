package userrepository

const (
	queryGetUserInfo = `
	SELECT vkid, name, true 
	FROM medics 
	WHERE vkid = $1 
	UNION ALL 
	SELECT vkid, name, false 
	FROM patients 
	WHERE vkid = $1;
	`
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
