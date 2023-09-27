
-- +migrate Up
CREATE TABLE "User" (
    id text PRIMARY KEY,
    name text
);

CREATE TABLE Transaction (
    id SERIAL,
    title text,
    price integer,
    user_id text,
    CONSTRAINT FK_transaction_user FOREIGN KEY (user_id) REFERENCES "User" (id)
);



-- +migrate Down