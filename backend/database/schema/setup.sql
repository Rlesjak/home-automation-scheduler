-- Enables the use of uuid_generate_v4() function
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Creates global id sequence, row ids are unique across the database
create SEQUENCE global_id_sequence