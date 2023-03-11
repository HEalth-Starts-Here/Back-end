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
    diaries (category, medicId, patientId, creatingDate, name, description)
VALUES
    (
		5,
        111,
        11,
        '2022-04-10 15:47:24',
        'Название дневника 1',
        'Описание дневника'
    );

INSERT INTO
    records (diaryId, posterPath, description)
VALUES
    (
        1,
        '/static/lesions/raw/1.png',
        'description text'
    );
