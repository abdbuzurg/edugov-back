CREATE TABLE IF NOT EXISTS employee_participation_in_professional_communities (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  professional_community_title VARCHAR(255) NOT NULL,
  language_code CHAR(2) NOT NULL,
  role_in_professional_community VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_participation_in_professional_communities_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_participation_in_professional_communities -- Corrected: Typo `communitites` -> `communities`
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_in_participation_in_professional_communities_language_code
  ON employee_participation_in_professional_communities (language_code);

