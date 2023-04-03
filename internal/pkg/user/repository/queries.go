package userrepository

const (
	queryMedicInit = `
	INSERT INTO
    medics (vkid, name)
	VALUES
    (
		$1,
        $2
    )
	RETURNING vkid, name;
	`

	queryPatientInit = `
	INSERT INTO
    patients (vkid, name)
	VALUES
    (
		$1,
        $2
    )
	RETURNING vkid, name;
	`
)
