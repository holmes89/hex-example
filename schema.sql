CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS tickets
(
  id uuid NOT NULL DEFAULT uuid_generate_v1(),
  creator varchar(255) NOT NULL,
  assigned varchar(255) DEFAULT NULL,
  title varchar(255) NOT NULL,
  description varchar(255) DEFAULT NULL,
  status varchar(255) NOT NULL,
  points integer DEFAULT NULL,
  created timestamp NOT NULL DEFAULT current_timestamp,
  updated timestamp NULL DEFAULT NULL,
  deleted timestamp NULL DEFAULT NULL
);
