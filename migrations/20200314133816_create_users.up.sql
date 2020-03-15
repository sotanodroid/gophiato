CREATE SCHEMA IF NOT EXISTS restapi;

CREATE TABLE IF NOT EXISTS restapi.users (
    id SERIAL NOT NULL PRIMARY KEY,
    email VARCHAR (50) NOT NULL,
    encrypted_password VARCHAR NOT NULL,
    UNIQUE(email)
);
