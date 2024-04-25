CREATE TABLE certification
(
    id                     UUID PRIMARY KEY,
    profile_id             UUID         NOT NULL REFERENCES profiles (id) ON DELETE CASCADE,
    name                   TEXT         NOT NULL,
    authority              TEXT         NOT NULL,
    license_number         TEXT,
    display_source         TEXT,
    url                    TEXT,
    authority_linkedin_url TEXT,
    start_date             TIMESTAMP(0),
    end_date               TIMESTAMP(0),
    created_at             TIMESTAMP(0) NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMP(0) NOT NULL DEFAULT NOW()
);