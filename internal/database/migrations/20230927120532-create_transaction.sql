
-- +migrate Up
CREATE TABLE "users" (
    id text PRIMARY KEY,
    name text
);

CREATE TABLE Transactions (
    id SERIAL PRIMARY KEY,
    title text,
    price integer,
    user_id text,
    CONSTRAINT FK_transactions_users FOREIGN KEY (user_id) REFERENCES "users" (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE Transactions;
DROP TABLE "users";
