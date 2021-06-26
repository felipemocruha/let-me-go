CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS entity (
  entity_id uuid DEFAULT uuid_generate_v4()
);
