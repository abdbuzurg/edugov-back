ALTER TABLE institution_details
  DROP institution_title,
  ADD COLUMN institution_title_short VARCHAR(255) NOT NULL,
  ADD COLUMN institution_title_long VARCHAR(255) NOT NULL,
  ALTER COLUMN legal_status DROP NOT NULL,
  ALTER COLUMN mission DROP NOT NULL,
  ALTER COLUMN founder DROP NOT NULL,
  ADD COLUMN city VARCHAR(255);
