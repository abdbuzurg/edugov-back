ALTER TABLE institution_details
  ADD COLUMN instititution_title VARCHAR(255) NOT NULL,
  DROP COLUMN institution_title_long,
  DROP COLUMN institution_title_short,
  ALTER COLUMN legal_status SET NOT NULL,
  ALTER COLUMN mission SET NOT NULL,
  ALTER COLUMN founder SET NOT NULL,
  DROP COLUMN city;

