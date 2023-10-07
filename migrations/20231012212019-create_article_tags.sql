
-- +migrate Up
CREATE TABLE article_tags (
    id bigserial PRIMARY KEY,
    article_id bigint NOT NULL,
    tag_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY article_tags
    ADD CONSTRAINT article_tags_article_id_tag_id_key UNIQUE (article_id, tag_id);
ALTER TABLE ONLY article_tags
    ADD CONSTRAINT article_tags_article_id_fkey FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE;
ALTER TABLE ONLY article_tags
    ADD CONSTRAINT article_tags_tag_id_fkey FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE CASCADE;

-- +migrate Down
DROP TABLE article_tags;
