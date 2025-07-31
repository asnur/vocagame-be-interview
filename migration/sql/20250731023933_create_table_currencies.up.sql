CREATE TABLE currencies (
    id SERIAL PRIMARY KEY,
    code VARCHAR(10) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

INSERT INTO currencies (code, name) VALUES
  ('USD', 'United States Dollar'),
  ('EUR', 'Euro'),
  ('JPY', 'Japanese Yen'),
  ('IDR', 'Indonesian Rupiah');
