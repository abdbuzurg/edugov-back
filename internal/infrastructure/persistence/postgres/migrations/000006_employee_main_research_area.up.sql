CREATE TABLE IF NOT EXISTS employee_main_research_areas (
  id BIGSERIAL,
  employee_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  area VARCHAR(255) NOT NULL,
  discipline VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_main_research_areas_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_employees_employee_main_research_areas
    FOREIGN KEY (employee_id) 
    REFERENCES employees (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_employee_main_research_areas_language_code
  ON employee_main_research_areas (language_code);

CREATE TABLE IF NOT EXISTS employee_main_research_area_key_topics (
  id BIGSERIAL,
  employee_main_research_area_id BIGINT NOT NULL,
  key_topic_title VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_main_research_area_key_topics_pkey
    PRIMARY KEY (id)
);
