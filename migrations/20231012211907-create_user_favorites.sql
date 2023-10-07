
-- +migrate Up
CREATE TABLE user_favorites (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    article_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY user_favorites
    ADD CONSTRAINT user_favorites_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE ONLY user_favorites
    ADD CONSTRAINT user_favorites_article_id_fkey FOREIGN KEY (article_id) REFERENCES articles (id) ON DELETE CASCADE;

-- +migrate Down
DROP TABLE user_favorites;
