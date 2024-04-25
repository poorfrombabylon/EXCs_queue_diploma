ALTER TABLE education DROP CONSTRAINT education_profile_id_key;

ALTER TABLE education
ADD COLUMN field_of_study TEXT,
ADD COLUMN degree_name TEXT,
ADD COLUMN school TEXT NOT NULL default '',
ADD COLUMN school_linkedin_profile_url TEXT,
ADD COLUMN description TEXT,
ADD COLUMN logo_url TEXT,
ADD COLUMN grade TEXT,
ADD COLUMN activities_and_societies TEXT,
ADD COLUMN start_date TIMESTAMP(0),
ADD COLUMN end_date TIMESTAMP(0);
