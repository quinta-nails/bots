-- migrate:up

CREATE TABLE bots (
    "id" BIGSERIAL PRIMARY KEY,
    "studio_id" BIGSERIAL NOT NULL UNIQUE,
    "token" VARCHAR NOT NULL UNIQUE,
    "first_name" VARCHAR NOT NULL,
    "username" VARCHAR NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

-- migrate:down

DROP TABLE bots;
