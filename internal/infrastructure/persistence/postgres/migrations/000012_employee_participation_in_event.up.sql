CREATE TABLE IF NOT EXISTS employee_participation_in_events (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  event_title VARCHAR(255) NOT NULL,
  event_date DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_participation_in_events_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_participation_in_events
    FOREIGN KEY (employee_id) -- Corrected: Added parentheses
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_participation_in_events_language_code
  ON employee_participation_in_events (language_code);
