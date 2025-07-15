CREATE TABLE IF NOT EXISTS institutions (
  id BIGSERIAL,
  year_of_establishment INT NOT NULL,
  email VARCHAR(127) NOT NULL,
  fax VARCHAR(31) NOT NULL,
  official_website VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institutions_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS institution_details (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  institution_type VARCHAR(127) NOT NULL,
  institution_title VARCHAR(255) NOT NULL,
  legal_status VARCHAR(127) NOT NULL,
  mission TEXT NOT NULL,
  founder VARCHAR(255) NOT NULL,
  legal_address VARCHAR(255) NOT NULL,
  factual_address VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institutions_details_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_details 
    FOREIGN KEY (institution_id) 
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_details_language_code 
  ON institution_details (language_code);

CREATE TABLE IF NOT EXISTS institution_accreditations (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  accreditation_type VARCHAR(127) NOT NULL,
  given_by VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_accreditations_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_accreditations 
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_accreditations_language_code 
  ON institution_accreditations (language_code);

CREATE TABLE IF NOT EXISTS institution_licences (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  licence_title VARCHAR(255) NOT NULL,
  licence_type VARCHAR(127) NOT NULL,
  link_to_file VARCHAR(255) NOT NULL,
  given_by VARCHAR(255) NOT NULL,
  date_start DATE NOT NULL,
  date_end DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_licences_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutons_institution_licences
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_licences_language_code
  ON institution_licences (language_code);

CREATE TABLE IF NOT EXISTS institution_rankings (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  ranking_title VARCHAR(127) NOT NULL,
  ranking_type VARCHAR(127) NOT NULL,
  date_recieved DATE NOT NULL,
  ranking_agency VARCHAR(100) NOT NULL,
  link_to_ranking_file VARCHAR(255) NOT NULL,
  description TEXT,
  link_to_ranking_agency VARCHAR(511) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_rankings_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_rankings
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_rankings_language_code
  ON institution_rankings (language_code);

CREATE TABLE IF NOT EXISTS institution_patents (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  patent_title VARCHAR(255) NOT NULL,
  discipline VARCHAR(128) NOT NULL,
  description TEXT NOT NULL,
  implemented_in TEXT NOT NULL,
  link_to_patent_file VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_patents_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_instituion_patents
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_patents_language_code 
  ON institution_patents (language_code);

CREATE TABLE IF NOT EXISTS institution_partnerships (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  partner_name VARCHAR(255) NOT NULL,
  partner_type VARCHAR(100) NOT NULL,
  date_of_contract DATE NOT NULL,
  link_to_partner VARCHAR(511) NOT NULL,
  goal TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_partnerships_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_partnerships
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_partnerships_language_code 
  ON institution_partnerships (language_code);

CREATE TABLE IF NOT EXISTS institution_socials (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  link_to_social VARCHAR(512) NOT NULL,
  social_name VARCHAR(64) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_socials_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_socials
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS institution_achievements (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  achievement_title VARCHAR(255) NOT NULL,
  achievement_type VARCHAR(127) NOT NULL,
  date_recieved DATE NOT NULL,
  given_by VARCHAR(255) NOT NULL,
  link_to_file VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_achievements_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_achievements
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_achievements_language_code
  ON institution_achievements (language_code);

CREATE TABLE IF NOT EXISTS institution_magazines (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  magazine_name VARCHAR(255) NOT NULL,
  link VARCHAR(512) NOT NULL,
  link_to_rinc VARCHAR(512),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  
  CONSTRAINT institution_magazines_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_magazines
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_magazines_language_code 
  ON institution_magazines (language_code);

CREATE TABLE IF NOT EXISTS institution_main_research_directions (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  research_direction_title VARCHAR(255) NOT NULL,
  discipline VARCHAR(127) NOT NULL,
  area_of_research VARCHAR(255),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  
  CONSTRAINT institution_main_research_directions_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_insitution_main_research_directions
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_main_research_directions_language_code 
  ON institution_main_research_directions (language_code);

CREATE TABLE IF NOT EXISTS institution_research_support_infrastructures (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  research_support_infrastructure_title VARCHAR(127) NOT NULL,
  research_support_infrastructure_type VARCHAR(127) NOT NULL,
  tin_of_legal_entity VARCHAR(15) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  
  CONSTRAINT institution_research_support_infrastructures_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_research_support_infrastructures
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_research_support_infrastructure_language_code
  ON institution_research_support_infrastructures (language_code);

CREATE TABLE IF NOT EXISTS institution_conferences (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  conference_title VARCHAR(255) NOT NULL,
  link VARCHAR(511) NOT NULL,
  link_to_rinc VARCHAR(511),
  date_of_conference DATE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_conferences_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_conferences
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_conferences_language_code 
  ON institution_conferences (language_code);

CREATE TABLE IF NOT EXISTS institution_projects (
  id BIGSERIAL,
  institution_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  project_type VARCHAR(15) NOT NULL,
  project_title VARCHAR(255) NOT NULL,
  date_start DATE NOT NULL,
  date_end DATE NOT NULL,
  fund DOUBLE PRECISION NOT NULL,
  institution_role TEXT NOT NULL,
  coordinator VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_projects_pkey 
    PRIMARY KEY (id),
  CONSTRAINT fk_institutions_institution_projects
    FOREIGN KEY (institution_id)
    REFERENCES institutions (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_projects_language_code
  ON institution_projects (language_code);

CREATE TABLE IF NOT EXISTS institution_project_partners (
  id BIGSERIAL,
  institution_project_id BIGINT NOT NULL,
  language_code CHAR(2) NOT NULL,
  partner_type VARCHAR(63) NOT NULL,
  partner_name VARCHAR(255) NOT NULL,
  link_to_partner VARCHAR(511) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT institution_project_partners_pkey
    PRIMARY KEY (id),
  CONSTRAINT fk_institution_projects_institution_project_partners
    FOREIGN KEY (institution_project_id)
    REFERENCES institution_projects (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_institution_project_partners_language_code
  ON institution_project_partners (language_code);
