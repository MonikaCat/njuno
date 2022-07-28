/* ---- TOKEN ---- */
CREATE TABLE token
(
    name TEXT NOT NULL UNIQUE
);


/* ---- TOKEN UNIT ---- */
CREATE TABLE token_unit
(
    token_name TEXT NOT NULL REFERENCES token (name),
    denom      TEXT NOT NULL UNIQUE,
    exponent   INT  NOT NULL,
    aliases    TEXT[],
    price_id   TEXT
);
