CREATE TABLE IF NOT EXISTS employee_refresher_courses (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  course_title VARCHAR(255) NOT NULL,
  date_start DATE NOT NULL,
  date_end DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_refresher_courses_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_refresher_courses
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_refresher_courses_language_code
  ON employee_refresher_courses (language_code);
