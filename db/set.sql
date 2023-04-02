DROP TABLE IF EXISTS medics               CASCADE;
DROP TABLE IF EXISTS patients             CASCADE;
DROP TABLE IF EXISTS diaries              CASCADE;
DROP TABLE IF EXISTS medicRecords         CASCADE;
DROP TABLE IF EXISTS patientRecords       CASCADE;
DROP TABLE IF EXISTS images               CASCADE;
DROP TABLE IF EXISTS tags                 CASCADE;


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
    patientId                           BIGINT,
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
    details                             VARCHAR(3000)
);

CREATE TABLE images (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    -- TODO add foreign key refer to patientRecords
    recordId                            BIGINT REFERENCES medicRecords (id) ON DELETE CASCADE,
    name                                VARCHAR(200)
);

CREATE TABLE tags (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    imageId                             BIGINT REFERENCES images (id) ON DELETE CASCADE,
    name                                VARCHAR(50)
);
