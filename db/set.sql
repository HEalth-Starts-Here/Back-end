DROP TABLE IF EXISTS diaries              CASCADE;
DROP TABLE IF EXISTS medics               CASCADE;
DROP TABLE IF EXISTS patients             CASCADE;
DROP TABLE IF EXISTS records             CASCADE;


CREATE TABLE medics (
    vkId                                BIGINT NOT NULL PRIMARY KEY
);

CREATE TABLE patients (
    vkId                                BIGINT NOT NULL PRIMARY KEY
);

CREATE TABLE diaries (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    category                            BIGINT NOT NULL,
    medicId                             BIGINT REFERENCES medics (vkId) ON DELETE CASCADE,
    patientId                           BIGINT REFERENCES patients (vkId) ON DELETE CASCADE,
    creatingDate                        TIMESTAMP NOT NULL,
    name                                VARCHAR(30) NOT NULL,
    description                         VARCHAR(3000)
);

CREATE TABLE records (
    id                                  BIGSERIAL NOT NULL PRIMARY KEY,
    diaryId                             BIGINT REFERENCES diaries (id) ON DELETE CASCADE,
    description                         VARCHAR(3000),
    posterPath                          VARCHAR(50) 
);
