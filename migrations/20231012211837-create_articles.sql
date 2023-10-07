
-- +migrate Up
CREATE TABLE articles (
    id bigserial PRIMARY KEY,
    author_id bigint NOT NULL,
    slug varchar(255) NULL,
    title varchar(255) NULL,
    description varchar(255) NOT NULL,
    body varchar(4096) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY articles
    ADD CONSTRAINT articles_slug_key UNIQUE (slug);
ALTER TABLE ONLY articles
    ADD CONSTRAINT articles_author_id_fkey FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE;

-- +migrate Down
DROP TABLE articles;
