/* ---- INFLATION ---- */
CREATE TABLE inflation
(
    one_row_id bool PRIMARY KEY DEFAULT TRUE,
    value      TEXT NOT NULL,
    height     BIGINT  NOT NULL,
    CONSTRAINT one_row_uni CHECK (one_row_id)
);
CREATE INDEX inflation_height_index ON inflation (height);


/* ---- STAKING POOL ---- */
CREATE TABLE staking_pool
(
    one_row_id        BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    bonded_tokens     TEXT    NOT NULL,
    not_bonded_tokens TEXT    NOT NULL,
    height            BIGINT  NOT NULL,
    CHECK (one_row_id)
);
CREATE INDEX staking_pool_height_index ON staking_pool (height);