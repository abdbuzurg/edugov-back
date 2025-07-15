CREATE TABLE IF NOT EXISTS employee_publications (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  publication_title VARCHAR(255) NOT NULL,
  link_to_publication VARCHAR(511) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_publications_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_publications
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_publications_language_code
  ON employee_publications (language_code);
