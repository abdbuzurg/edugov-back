CREATE TABLE IF NOT EXISTS employee_details (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  surname VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  middlename VARCHAR(255) NOT NULL,
  is_employee_details_new BOOLEAN NOT NULL DEFAULT false,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_details_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_details
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_details_language_code
  ON employee_details (language_code);
