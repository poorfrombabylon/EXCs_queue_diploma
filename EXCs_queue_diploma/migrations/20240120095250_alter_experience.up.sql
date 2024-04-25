ALTER TABLE experience DROP CONSTRAINT experience_profile_id_key;

ALTER TABLE experience
ADD COLUMN position TEXT,
ADD COLUMN company_name TEXT NOT NULL DEFAULT '',
ADD COLUMN location TEXT,
ADD COLUMN description TEXT,
ADD COLUMN start_date TIMESTAMP(0),
ADD COLUMN end_date TIMESTAMP(0);