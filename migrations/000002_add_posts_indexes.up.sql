CREATE INDEX IF NOT EXISTS posts_title_idx ON posts USING GIN (to_tsvector('simple', title));
CREATE INDEX IF NOT EXISTS posts_img_urls_idx ON posts USING GIN (img_urls);