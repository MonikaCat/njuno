
CREATE TABLE validator_info
(
    consensus_address     TEXT   NOT NULL UNIQUE PRIMARY KEY REFERENCES validator (consensus_address),
    operator_address      TEXT   NOT NULL UNIQUE,
    self_delegate_address TEXT REFERENCES account (address),
    height                BIGINT NOT NULL
);
CREATE INDEX validator_info_operator_address_index ON validator_info (operator_address);
CREATE INDEX validator_info_self_delegate_address_index ON validator_info (self_delegate_address);

CREATE TABLE validator_voting_power
(
    validator_address TEXT   NOT NULL REFERENCES validator (consensus_address) PRIMARY KEY,
    voting_power      BIGINT NOT NULL,
    height            BIGINT NOT NULL
);
CREATE INDEX validator_voting_power_height_index ON validator_voting_power (height);

CREATE TABLE validator_description
(
    validator_address TEXT   NOT NULL REFERENCES validator (consensus_address) PRIMARY KEY,
    moniker           TEXT,
    details           TEXT,
    height            BIGINT NOT NULL
);
CREATE INDEX validator_description_height_index ON validator_description (height);