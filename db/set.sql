DROP TABLE IF EXISTS diaries              CASCADE;
DROP TABLE IF EXISTS medics               CASCADE;
DROP TABLE IF EXISTS patients             CASCADE;
DROP TABLE IF EXISTS records              CASCADE;
DROP TABLE IF EXISTS images               CASCADE;


CREATE TABLE medics (
    vkId                                BIGINT NOT NULL PRIMARY KEY
);

CREATE TABLE patients (
    vkId                                BIGINT NOT NULL PRIMARY KEY
    -- TODO: add characterestics
);

CREATE TABLE diaries (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    category                            BIGINT NOT NULL,
    medicId                             BIGINT REFERENCES medics (vkId) ON DELETE CASCADE,
    patientId                           BIGINT REFERENCES patients (vkId) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    title                               VARCHAR(50) NOT NULL,
    description                         VARCHAR(3000)
);

CREATE TABLE records (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    diaryId                             BIGINT REFERENCES diaries (id) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    title                               VARCHAR(50),
    description                         VARCHAR(3000),
    area                                FLOAT DEFAULT 0.0,
    dryness                             SMALLINT DEFAULT 0 CHECK (dryness >= 0 AND dryness <= 10),
    edema                               SMALLINT DEFAULT 0 CHECK (edema >= 0 AND edema <= 10),
    itching                             SMALLINT DEFAULT 0 CHECK (itching >= 0 AND itching <= 10),
    pain                                SMALLINT DEFAULT 0 CHECK (pain >= 0 AND pain <= 10),
    peeling                             SMALLINT DEFAULT 0 CHECK (peeling >= 0 AND peeling <= 10),
    redness                             SMALLINT DEFAULT 0 CHECK (redness >= 0 AND redness <= 10)
);

CREATE TABLE images (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    recordId                            BIGINT REFERENCES records (id) ON DELETE CASCADE,
    name                                VARCHAR(200), 
    area                                FLOAT DEFAULT 0.0
);
