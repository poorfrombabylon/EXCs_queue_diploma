CREATE TABLE skills
(
    id          UUID PRIMARY KEY,
    profile_id UUID NOT NULL REFERENCES profiles (id) ON DELETE CASCADE,
    skills     json,
    languages  json,
    created_at TIMESTAMP(0) NOT NULL DEFAULT NOW()
);