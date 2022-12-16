CREATE TABLE IF NOT EXISTS categories (
  id BIGSERIAL,
  name VARCHAR(255),

  created_at TIMESTAMPTZ DEFAULT now(),
  updated_at TIMESTAMPTZ DEFAULT now(),

  PRIMARY KEY(id)
)