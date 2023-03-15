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
    records (diaryId, creatingDate, title, description)
VALUES
    (
        1,
        '2022-04-10 15:47:24',
        'Title records 1 of 1 diary',
        'Description records 1 of 1 diary'
    ),
    (
        1,
        '2022-04-10 15:47:24',
        'Title records 2 of 1 diary',
        'Description records 2 of 1 diary'
    ),
    (
        2,
        '2022-04-10 15:47:24',
        'Title records of 1 diary',
        'Description records of 1 diary'
    );
