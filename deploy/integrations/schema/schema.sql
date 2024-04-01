CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users  (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- Generate UUID v4
  full_name VARCHAR(255) NOT NULL,
  phone_number VARCHAR(20),
  email VARCHAR(255),
  user_name VARCHAR(50),
  password VARCHAR(255) NOT NULL,
  dob DATE,
  latest_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT check_register CHECK (phone_number IS NOT NULL OR email IS NOT NULL OR user_name IS NOT NULL)
);

