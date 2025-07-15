CREATE TABLE IF NOT EXISTS employee_patents (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL, 
  patent_title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_patents_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_patents
    FOREIGN KEY (employee_id)
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_patents_language_code
  ON employee_patents (language_code);
