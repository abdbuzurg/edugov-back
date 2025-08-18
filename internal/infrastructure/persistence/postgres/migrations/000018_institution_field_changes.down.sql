ALTER TABLE institutions
  ALTER COLUMN year_of_establishment SET NOT NULL,
  ALTER COLUMN fax SET NOT NULL,
  DROP COLUMN phone_number,
  DROP COLUMN mail_index;