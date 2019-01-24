CREATE TABLE IF NOT EXISTS stats (
    id VARCHAR(255)           NOT NULL,
    int1 INT            NOT NULL,
    int2 INT            NOT NULL,
    limit_range INT     NOT NULL,
    str1 VARCHAR(255)   NOT NULL,
    str2 VARCHAR(255)   NOT NULL,
    nb_record INT    NOT NULL,
    PRIMARY KEY (id)
)