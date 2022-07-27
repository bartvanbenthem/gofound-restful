CREATE TABLE IF NOT EXISTS posts (
id bigserial PRIMARY KEY,
created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
title text NOT NULL,
content text NOT NULL,
author text NOT NULL,
img_urls text[] NOT NULL,
version integer NOT NULL DEFAULT 1
);
