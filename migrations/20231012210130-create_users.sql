
-- +migrate Up
CREATE TABLE users (
    id bigserial PRIMARY KEY,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    bio varchar(4096),
    image varchar(4096),
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_username_email_key UNIQUE (username, email);

-- +migrate Down
DROP TABLE users;
