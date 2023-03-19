INSERT INTO
    medics (vkid)
VALUES
    (
        111
    ),
    (
        222
    );

INSERT INTO
    patients (vkid)
VALUES
    (
        11
    ),
    (
        22
    );

INSERT INTO
    diaries (category, medicId, patientId, creatingDate, title, description)
VALUES
    (
		5,
        111,
        11,
        '2022-04-10 15:47:24',
        'Название дневника 1',
        'Описание дневника 2'
    ),
    (
		1,
        111,
        11,
        '2022-04-10 15:47:24',
        'Название дневника 2',
        'Описание дневника 2'
    );

INSERT INTO
    records (diaryId, creatingDate, title, description, dryness, edema, itching, pain, peeling, redness)
VALUES
    (
        1,
        '2022-04-10 15:47:24',
        'Title records 1 of 1 diary',
        'Description records 1 of 1 diary',
        1,
        1,
        1,
        1,
        1,
        1
    ),
    (
        1,
        '2022-04-10 15:47:24',
        'Title records 2 of 1 diary',
        'Description records 2 of 1 diary',
        1,
        1,
        1,
        1,
        1,
        1
    ),
    (
        2,
        '2022-04-10 15:47:24',
        'Title records of 1 diary',
        'Description records of 1 diary',
        1,
        1,
        1,
        1,
        1,
        1
    );

INSERT INTO
    images (recordId, name)
VALUES
    (
		1,
        '1.png'
    ),
    (
		1,
        '2.png'
    ),
    (
		1,
        '3.png'
    ),
    (
		2,
        '4.png'
    ),
    (
		2,
        '5.png'
    );
