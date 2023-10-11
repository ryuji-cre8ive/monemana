
-- +migrate Up
CREATE TABLE "users" (
    id text PRIMARY KEY,
    name text
);

CREATE TABLE Transactions (
    id text PRIMARY KEY,
    title text,
    price decimal,
    user_id text,
    date timestamp,
    rate decimal,
    CONSTRAINT FK_transactions_users FOREIGN KEY (user_id) REFERENCES "users" (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE Transactions;
DROP TABLE "users";
