
-- +migrate Up
CREATE TABLE tags (
    id bigserial PRIMARY KEY,
    description varchar(255) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY tags
    ADD CONSTRAINT tags_description_key UNIQUE (description);

-- +migrate Down
DROP TABLE tags;
