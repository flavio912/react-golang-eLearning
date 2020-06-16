-- Allows for the uuid_generate_v4() command
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER ROLE ttc-api SUPERUSER;