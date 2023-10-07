
-- +migrate Up
CREATE TABLE follows (
    id bigserial PRIMARY KEY,
    follower_id bigint NOT NULL,
    following_id bigint NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE ONLY follows
    ADD CONSTRAINT follows_follower_id_fkey FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE ONLY follows
    ADD CONSTRAINT follows_following_id_fkey FOREIGN KEY (following_id) REFERENCES users (id) ON DELETE CASCADE;

-- +migrate Down
DROP TABLE follows;
