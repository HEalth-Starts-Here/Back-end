INSERT INTO
    medics (vkid, name)
VALUES
    (
        111,
        'Иван Иванов'
    ),
    (
        222,
        'Петр Петров'
    );

INSERT INTO
    patients (vkid, name)
VALUES
    (
        11,
        'Максим Максимов'
    ),
    (
        22,
        'Александр Александров'
    );

INSERT INTO
    diaries (medicId, patientId, creatingDate, title, complaints, anamnesis, objectively, diagnosis)
VALUES
    (
        111,
        11,
        '2022-04-10 15:47:24',
        'Название дневника 1',
        'Жалобы дневника 1',
        'Анамнез дневника 1',
        'Объективно дневника 1',
        'Диагноз дневника 1'
    ),
    (
        111,
        11,
        '2022-04-10 15:47:24',
        'Название дневника 2',
        'Жалобы дневника 2',
        'Анамнез дневника 2',
        'Объективно дневника 2',
        'Диагноз дневника 2'
    ),
    (
        111,
        NULL,
        '2022-04-10 15:47:24',
        'Название дневника 2',
        'Жалобы дневника 2',
        'Анамнез дневника 2',
        'Объективно дневника 2',
        'Диагноз дневника 2'
    );

INSERT INTO
    medicRecords (diaryId, creatingDate, title, treatment, recommendations, details)
VALUES
    (
        1,
        '2022-04-10 15:47:24',
        'Название записи врача 1 дневника 1',
        'Лечение в записи врача 1 дневника 1',
        'Рекомендации в записи врача 1 дневника 1',
        'Подробности в записи врача 1 дневника 1'
        -- 'diarisation'
    ),
    (
        1,
        '2022-04-10 15:47:24',
        'Название записи врача 2 дневника 1',
        'Лечение в записи врача 2 дневника 1',
        'Рекомендации в записи врача 2 дневника 1',
        'Подробности в записи врача 2 дневника 1'
        -- 'diarisation'
    ),
    (
        2,
        '2022-04-10 15:47:24',
        'Название записи врача 1 дневника 2',
        'Лечение в записи врача 1 дневника 2',
        'Рекомендации в записи врача 1 дневника 2',
        'Подробности в записи врача 1 дневника 2'
        -- 'diarisation'
    );

INSERT INTO
    patientRecords (diaryId, creatingDate, title, complaints, treatment, details, feelings)
VALUES
    (
        1,
        '2022-04-10 15:47:24',
        'Название записи пациента 1 дневника 1',
        'Жалобы в записи пациента 1 дневника 1',
        'Лечение в записи пациента 1 дневника 1',
        'Подробности в записи пациента 1 дневника 1',
        6
    ),
    (
        1,
        '2022-04-10 15:47:24',
        'Название записи пациента 2 дневника 1',
        'Жалобы в записи пациента 2 дневника 1',
        'Лечение в записи пациента 2 дневника 1',
        'Подробности в записи пациента 2 дневника 1',
        8
    ),
    (
        2,
        '2022-04-10 15:47:24',
        'Название записи пациента 1 дневника 2',
        'Жалобы в записи пациента 1 дневника 2',
        'Лечение в записи пациента 1 дневника 2',
        'Подробности в записи пациента 1 дневника 2',
        8
    );

INSERT INTO
    -- images (ismedic, recordId, name)
    medicImages (recordId, name)
VALUES
    (
        -- true,
		1,
        '1.png'
    ),
    (
        -- true,
		1,
        '2.png'
    ),
    (
        -- true,
		1,
        '3.png'
    ),
    (
        -- false,
		2,
        '4.png'
    ),
    (
        -- false,
		2,
        '5.png'
    );


INSERT INTO
    -- patientImages (ismedic, recordId, name)
    patientImages ( recordId, name)
VALUES
    (
        -- false,
		1,
        '1.png'
    ),
    (
        -- false,
		1,
        '2.png'
    ),
    (
        -- false,
		1,
        '3.png'
    ),
    (
        -- false,
		2,
        '4.png'
    ),
    (
        -- false,
		2,
        '5.png'
    );


INSERT INTO
    audio (medicRecordId, creatingDate, diarisation, filename)
VALUES
    (
        1,
        '2022-04-10 15:47:24',
		'Diarisation 1 of medic record 1',
        '1.mp3'
    ),
    (
        1,
        '2022-04-10 15:47:24',
		'Diarisation 2 of medic record 1',
        '2.mp3'
    ),
    (
        1,
        '2022-04-10 15:47:24',
		'Diarisation 3 of medic record 1',
        '3.mp3'
    ),
    (
        2,
        '2022-04-10 15:47:24',
		'Diarisation 1 of medic record 2',
        '4.mp3'
    ),
    (
        2,
        '2022-04-10 15:47:24',
		'Diarisation 2 of medic record 2',
        '5.mp3'
    );

INSERT INTO
    comments (authorIsMedic, creatingDate, diaryId, text)
VALUES
    (
        false,
        '2022-04-10 15:47:24',
		1,
        'Text of comment 1 of diary 1'
    ),
    (
        true,
        '2022-04-10 15:47:24',
		1,
        'Text of comment 2 of diary 1'
    ),
    (
        false,
        '2022-04-10 15:47:24',
		1,
        'Text of comment 2 of diary 1'
    );

INSERT INTO
    notes (medicRecordId, patientRecordId, IsMedicRecord, creatingDate, text)
VALUES
    (
        1,
        NULL,
        true,
        '2022-04-10 15:47:24',
        'Text of note 1 of medic record 1'
    ),
    (
        1,
        NULL,
        true,
        '2022-04-10 15:47:24',
        'Text of note 2 of medic record 1'
    ),
    (
        2,
        NULL,
        true,
        '2022-04-10 15:47:24',
        'Text of note 1 of medic record 2'
    ),
    (
        NULL,
        1,
        false,
        '2022-04-10 15:47:24',
        'Text of note 1 of patient record 1'
    ),
    (
        NULL,
        1,
        false,
        '2022-04-10 15:47:24',
        'Text of note 2 of patient record 1'
    );

INSERT INTO
    diarytokens (diaryid, token)
VALUES
    (
        3,
        'linktoken'
    );
    
-- CREATE TABLE notes (
--     medicRecordId                       BIGINT REFERENCES medicRecords (id) ON DELETE CASCADE,
--     patientRecordId                     BIGINT REFERENCES patientRecords (id) ON DELETE CASCADE,
--     recordAuthorIsMedic                 BOOLEAN,
--     id                                  BIGSERIAL NOT NULL PRIMARY KEY,
--     creatingDate                        TIMESTAMP NOT NULL,
--     text                                VARCHAR(1000)
-- );

-- INSERT INTO
--     records_notes (medicRecordId, patientRecordId, recordAuthorIsMedic, noteId)
-- VALUES
--     (

--         '2022-04-10 15:47:24',
--         'Text of note 1 of record 1'
--     ),
--     (
--         '2022-04-10 15:47:24',
--         'Text of note 2 of record 1'
--     ),
--     (
--         '2022-04-10 15:47:24',
--         'Text of note 1 of record 2'
--     ),;

-- CREATE TABLE records_notes (
--     medicRecordId                       BIGINT REFERENCES medicRecords (id) ON DELETE CASCADE,
--     patientRecordId                     BIGINT REFERENCES patientRecords (id) ON DELETE CASCADE,
--     recordAuthorIsMedic                 BOOLEAN,
--     noteId                              BIGINT REFERENCES notes (id) ON DELETE CASCADE
-- );
-- INSERT INTO
--     tags (imageId, name)
-- VALUES
--     (
-- 		1,
--         'левая нога'
--     ),
--     (
-- 		1,
--         'правая рука'
--     ),
--     (
-- 		1,
--         'лицо'
--     ),
--     (
-- 		2,
--         'нос'
--     ),
--     (
-- 		2,
--         'грудь'
--     );