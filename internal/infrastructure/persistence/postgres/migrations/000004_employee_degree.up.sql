CREATE TABLE IF NOT EXISTS employee_degrees (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  university_name VARCHAR(255) NOT NULL,
  degree_level VARCHAR(127) NOT NULL,
  speciality VARCHAR(255) NOT NULL,
  date_start DATE NOT NULL,
  date_end DATE NOT NULL,
  given_by VARCHAR(255),
  date_degree_recieved DATE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_degrees_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_degrees
    FOREIGN KEY (employee_id) -- Corrected: Added parentheses
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_degrees_language_code
  ON employee_degrees (language_code);
