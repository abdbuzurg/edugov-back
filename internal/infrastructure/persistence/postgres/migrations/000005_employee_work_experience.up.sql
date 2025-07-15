CREATE TABLE IF NOT EXISTS employee_work_experiences (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  workplace VARCHAR(255) NOT NULL,
  job_title VARCHAR(255) NOT NULL, 
  description VARCHAR(255) NOT NULL,
  date_start DATE NOT NULL,
  date_end DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_work_experience_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_work_experience
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_work_experience_language_code
  ON employee_work_experiences (language_code);
