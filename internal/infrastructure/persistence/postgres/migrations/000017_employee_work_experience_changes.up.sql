ALTER TABLE employee_work_experiences 
  ALTER COLUMN date_end TYPE DATE, 
  ALTER COLUMN date_end DROP NOT NULL;

ALTER TABLE employee_work_experiences
  ADD COLUMN on_going BOOLEAN NOT NULL DEFAULT FALSE;