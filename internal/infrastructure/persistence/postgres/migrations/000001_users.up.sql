CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  user_type VARCHAR(50) NOT NULL,
  entity_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT users_pkey
    PRIMARY KEY (id)
);

CREATE INDEX IF NOT EXISTS idx_users_user_type ON users(user_type);

CREATE TABLE IF NOT EXISTS user_session (
  id BIGSERIAL,
  user_id BIGINT NOT NULL,
  refresh_token TEXT UNIQUE NOT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

  CONSTRAINT user_session_pkey
    PRIMARY KEY (id),

  CONSTRAINT fk_users_user_session
    FOREIGN KEY (user_id)
    REFERENCES users (id)
    ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_user_session_refresh_token ON user_session(refresh_token);
