ALTER TABLE employee_work_experiences 
ALTER COLUMN date_end TYPE DATE, 
ALTER COLUMN date_end SET NOT NULL;

ALTER TABLE employee_work_experiences
DROP COLUMN on_going;