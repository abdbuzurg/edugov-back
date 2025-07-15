CREATE TABLE IF NOT EXISTS employees (
  id BIGSERIAL,
  unique_id VARCHAR(9) UNIQUE NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT employee_pkey
    PRIMARY KEY (id)
);
