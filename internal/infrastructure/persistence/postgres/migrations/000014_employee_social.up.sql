CREATE TABLE IF NOT EXISTS employee_socials (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  social_name VARCHAR(31) NOT NULL,
  link_to_social VARCHAR(511) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_socials_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_socials
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);
