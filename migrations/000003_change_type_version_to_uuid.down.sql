-- Add a new column to the movies table
ALTER TABLE movies
    ADD COLUMN version_int integer;

-- Drop the old version column
ALTER TABLE movies
    DROP COLUMN version;

-- Rename the new version_uuid column to version
ALTER TABLE movies
    RENAME COLUMN version_int TO version;