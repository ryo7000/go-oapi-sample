
-- +migrate Up
CREATE TABLE comments (
    id bigserial PRIMARY KEY,
    author_id bigint NOT NULL,
    article_id bigint NOT NULL,
    body varchar(4096) NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY comments
    ADD CONSTRAINT comments_author_id_fkey FOREIGN KEY (author_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE ONLY comments
    ADD CONSTRAINT comments_article_id_fkey FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE;

-- +migrate Down
DROP TABLE comments;
