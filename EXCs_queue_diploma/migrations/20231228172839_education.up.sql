CREATE TABLE education
(
    id          UUID PRIMARY KEY,
    profile_id UUID NOT NULL REFERENCES profiles (id) ON DELETE CASCADE UNIQUE,
    education json,
    created_at TIMESTAMP(0) NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) NOT NULL DEFAULT NOW()
);