
-- +migrate Up
CREATE TABLE "users" (
    id text PRIMARY KEY,
    display_name text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE rooms (
    id text PRIMARY KEY,
    display_name text,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE rooms_binding_users (
    id text PRIMARY KEY,
    room_id text,
    user_id text,
    CONSTRAINT FK_room_binding_users_users FOREIGN KEY (user_id) REFERENCES "users" (id) ON DELETE CASCADE,
    CONSTRAINT FK_room_binding_users_rooms FOREIGN KEY (room_id) REFERENCES "rooms" (id) ON DELETE CASCADE
);

CREATE TABLE transactions (
    id text PRIMARY KEY,
    title text,
    price decimal,
    user_id text,
    target_user_id text,
    room_id text,
    created_at timestamp,
    deleted_at timestamp,
    CONSTRAINT FK_transactions_users FOREIGN KEY (user_id) REFERENCES "users" (id) ON DELETE CASCADE,
    CONSTRAINT FK_transactions_target_users FOREIGN KEY (target_user_id) REFERENCES "users" (id) ON DELETE CASCADE,
    CONSTRAINT FK_transactions_rooms FOREIGN KEY (room_id) REFERENCES "rooms" (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE transactions;
DROP TABLE rooms_binding_users;
DROP TABLE rooms;
DROP TABLE "users";

