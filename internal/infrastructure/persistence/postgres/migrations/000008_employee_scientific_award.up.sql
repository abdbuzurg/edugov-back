CREATE TABLE IF NOT EXISTS employee_scientific_awards (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  scientific_award_title VARCHAR(255) NOT NULL,
  given_by VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_scientific_awards_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_scientific_awards
    FOREIGN KEY (employee_id) -- Corrected: Added parentheses
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_scientific_awards_language_code
  ON employee_scientific_awards (language_code);
