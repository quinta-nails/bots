-- migrate:up

CREATE TABLE telegram_users (
    "user_id" BIGINT NOT NULL UNIQUE,
    "telegram_id" BIGINT NOT NULL UNIQUE,
    "username" VARCHAR,
    "created_at" timestamptz NOT NULL DEFAULT NOW()
);

-- migrate:down

DROP TABLE telegram_users;

