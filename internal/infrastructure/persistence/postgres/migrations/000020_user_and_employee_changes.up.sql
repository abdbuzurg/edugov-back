ALTER TABLE employees 
  ADD COLUMN user_id BIGINT;

UPDATE employees
SET user_id = (
  SELECT id
  FROM users
  WHERE users.entity_id = employees.id AND users.user_type = 'employee'
);

ALTER TABLE employees
  ADD CONSTRAINT fk_employees_user
  FOREIGN KEY (user_id) REFERENCES users (id);

CREATE INDEX idx_employee_user_id ON employees(user_id);

ALTER TABLE users
  DROP COLUMN entity_id,
  DROP COLUMN user_type;
