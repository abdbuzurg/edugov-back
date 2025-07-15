CREATE TABLE IF NOT EXISTS employee_research_activities (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  research_activity_title VARCHAR(255) NOT NULL,
  employee_role TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_research_activities_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_research_activities
    FOREIGN KEY (employee_id) -- Corrected: Added parentheses
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_research_activities_language_code
  ON employee_research_activities (language_code);

