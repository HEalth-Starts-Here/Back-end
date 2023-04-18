DROP TABLE IF EXISTS medics               CASCADE;
DROP TABLE IF EXISTS patients             CASCADE;
DROP TABLE IF EXISTS diaries              CASCADE;
DROP TABLE IF EXISTS medicRecords         CASCADE;
DROP TABLE IF EXISTS patientRecords       CASCADE;
DROP TABLE IF EXISTS medicImages          CASCADE;
DROP TABLE IF EXISTS patientImages        CASCADE;
DROP TABLE IF EXISTS tags                 CASCADE;
DROP TABLE IF EXISTS audio                CASCADE;
DROP TABLE IF EXISTS comments             CASCADE;
DROP TABLE IF EXISTS records_notes        CASCADE;
DROP TABLE IF EXISTS notes                CASCADE;
DROP TABLE IF EXISTS diaryTokens          CASCADE;


CREATE TABLE medics (
    vkId                                BIGINT NOT NULL PRIMARY KEY,
    name                                VARCHAR(200)
);

CREATE TABLE patients (
    vkId                                BIGINT NOT NULL PRIMARY KEY,
    name                                VARCHAR(200)
    -- TODO: add characterestics
);

CREATE TABLE diaries (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    medicId                             BIGINT REFERENCES medics (vkId) ON DELETE CASCADE,
    -- TODO: add reference to patients table
    -- patientId                           BIGINT,
    patientId                           BIGINT REFERENCES patients (vkId) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    title                               VARCHAR(50) NOT NULL,
    complaints                          VARCHAR(1000),
    anamnesis                           VARCHAR(1000),
    objectively                         VARCHAR(1000),
    diagnosis                           VARCHAR(1000)
);

CREATE TABLE medicRecords (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    diaryId                             BIGINT REFERENCES diaries (id) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    title                               VARCHAR(50),
    treatment                           VARCHAR(1000),
    recommendations                     VARCHAR(1000),
    details                             VARCHAR(3000)
);

CREATE TABLE patientRecords (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    diaryId                             BIGINT REFERENCES diaries (id) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    title                               VARCHAR(50),
    complaints                          VARCHAR(1000),
    treatment                           VARCHAR(1000),
    details                             VARCHAR(3000),
    feelings                            BIGINT DEFAULT 0
);

CREATE TABLE medicImages (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    -- TODO add foreign key refer to patientRecords
    -- isMedic                             BOOLEAN,
    recordId                            BIGINT REFERENCES medicRecords (id) ON DELETE CASCADE,
    name                                VARCHAR(200)
);

CREATE TABLE patientImages (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    -- isMedic                             BOOLEAN,
    recordId                            BIGINT REFERENCES patientRecords (id) ON DELETE CASCADE,
    name                                VARCHAR(200)
);

-- CREATE TABLE tags (
--     id                                  BIGSERIAL NOT NULL PRIMARY KEY,
--     imageId                             BIGINT REFERENCES images (id) ON DELETE CASCADE,
--     name                                VARCHAR(50)
-- );

CREATE TABLE audio (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    medicRecordId                       BIGINT REFERENCES medicRecords (id) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    diarisation                         VARCHAR(10000),
    filename                            VARCHAR(200)
);

CREATE TABLE comments (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    authorIsMedic                       BOOLEAN,
    isreaded                            BOOLEAN DEFAULT false,
    creatingDate                        TIMESTAMP NOT NULL,
    diaryId                             BIGINT REFERENCES diaries (id) ON DELETE CASCADE,
    text                                VARCHAR(1000)
);

CREATE TABLE notes (
    medicRecordId                       BIGINT REFERENCES medicRecords (id) ON DELETE CASCADE,
    patientRecordId                     BIGINT REFERENCES patientRecords (id) ON DELETE CASCADE,
    IsMedicRecord                       BOOLEAN,
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    creatingDate                        TIMESTAMP NOT NULL,
    text                                VARCHAR(1000)
);

CREATE TABLE diaryTokens (
    diaryId                              BIGINT REFERENCES diaries (id) ON DELETE CASCADE,
    token                                VARCHAR(256)
);
