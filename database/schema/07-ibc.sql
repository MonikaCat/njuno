/* ---- IBC TRANSFER PARAMS ---- */
CREATE TABLE ibc_transfer_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    height     BIGINT  NOT NULL,
    CHECK (one_row_id)
);
CREATE INDEX ibc_transfer_params_height_index ON ibc_transfer_params (height);
