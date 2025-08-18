ALTER TABLE institutions
  ALTER COLUMN year_of_establishment DROP NOT NULL,
  ALTER COLUMN fax DROP NOT NULL,
  ADD COLUMN phone_number VARCHAR(255),
  ADD COLUMN mail_index VARCHAR(255);
  
