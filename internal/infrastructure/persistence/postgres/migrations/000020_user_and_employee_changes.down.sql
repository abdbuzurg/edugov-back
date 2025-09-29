ALTER TABLE users
  ADD COLUMN entity_id BIGINT,
  ADD COLUMN user_type VARCHAR(15);

UPDATE users
SET 
  user_type = 'employee',
  entity_id = (
    SELECT id
    FROM employees
    WHERE users.id = employees.user_id
  );

ALTER TABLE employees
  DROP COLUMN user_id;

DROP INDEX IF EXISTS idx_employee_user_id;
