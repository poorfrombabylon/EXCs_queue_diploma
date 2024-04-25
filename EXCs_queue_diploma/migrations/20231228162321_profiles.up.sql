CREATE TABLE profiles
(
    id              UUID PRIMARY KEY,
    first_name      TEXT         NOT NULL,
    last_name       TEXT         NOT NULL,
    country         TEXT,
    city            TEXT,
    state           TEXT,
    gender          TEXT,
    occupation      TEXT,
    summary         TEXT,
    linkedin_id     TEXT         NOT NULL UNIQUE,
    created_at      TIMESTAMP(0) NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP(0) NOT NULL DEFAULT NOW(),
    last_checked_at TIMESTAMP(0) NOT NULL DEFAULT NOW()
);