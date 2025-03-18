

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

COMMENT ON SCHEMA public IS 'standard public schema';

CREATE SCHEMA app;

ALTER SCHEMA app OWNER TO dbuser;


CREATE TABLE app.task (
    id varchar,
    name varchar,
    description varchar,
    creation_user_id character varying,
    update_user_id character varying,
    created_at timestamp without time zone,
    updated_at timestamp
);

ALTER TABLE app.task OWNER TO dbuser;

INSERT INTO app.task VALUES ('5fdd0823-3a3d-49b7-9452-ec994a03cca1','sdjib6jnIT','udDHjXOoiE','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('bb9068b1-4e56-4d0e-8f37-a78b544a9062','KcqJq1yXAY','M0lRJTumVn','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('d8fc69d0-036f-47ab-ae80-b502c4d0a0d2','UdpKZiyrz7','jkv4oUqhLF','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('8d43f7af-b7aa-478b-9ea5-6c4d7e9c6fc8','WEGQPz430D','oJyZghhjdt','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('435f8fa7-e7d9-4291-b795-4406aedee642','XR2u1eQDK5','xdVlUS870J','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('11e05ccc-fd83-4456-a7c8-8574f3a73aa7','nmMmNcsNHz','RsnAWfx14C','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('324fb597-7658-41bd-87c6-052c866698d4','HBFhoLxCW1','GFaclgfkSb','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('e7c3a7c0-2548-4baf-b8f2-8a96fcb9ad01','ES81HcyV2p','3Mn8dUD87t','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('a7abd06b-6362-471e-b1df-92bc54230826','1SdtAUF3Yd','alpzVcQqs4','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('c9fcc77d-ec51-419b-a1a2-41b7b901d8bd','tKrLwV7czg','ZTC1BlZWzH','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('b069ebb0-3183-45a4-8986-7719901a7e8c','lsqBzKjG7m','ltdfdfMSdU','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('6a954f67-0d90-4765-920a-14c3b78cb9a6','2h0wlQjkHb','6xx3BaSTUD','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('c886d592-8d7c-4546-ad9e-4e76578d7ce4','W9RS47zMNA','zcDrZQnH0I','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('6e703426-2b42-434a-b812-a3d978402e56','3DNlQyZMEV','RrSefnbgUq','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('0b5cae50-a79c-40e5-973c-0cccc319468a','48D9i0N3xU','kPf7KsFW6Y','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('2906cbec-c03c-4441-87b0-b3663598cf29','RXKkpxngKK','4t8YkRhpoA','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('f002472b-5446-43b0-917d-dc57504632f6','SqhbUVaXkj','JuiI3iwXGI','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('0bdf795a-d421-4982-9ec6-28fbf69f77b8','8Ge0qsQMIG','hf1v4zJzBH','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('9065fa84-44c0-44a9-af30-272eed409d80','TxHCwLQgIW','5Taz6d4wh0','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('3c07b9a0-6dca-4b86-88c5-b726db7ffe3d','GAboBdN4ah','968Pu9Whax','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('2ed6cda2-4fd7-4fb4-a7dd-c58d1a4ea9fd','9OnCYc83Ug','QDkcXwwpKa','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('e97447f7-80f2-495d-bcd6-9333604fb1f5','qkDZ7aQ8Rl','0i3ug2tsmy','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('72f7c59f-bc55-444e-a846-fa71f65843f5','LF9wDSOGnz','z8UWmxPT5y','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('df51d1c5-4b29-4e8a-b4b2-469445414ae0','tYy0h8OqIh','1uhVFj8vxB','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('941b633e-0c5d-44ed-a40b-fbf55499ed94','lOJDfz74Le','zETygMXKsY','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('ae6ec983-c01d-4ad6-9ee5-7202f811de61','2ebn8iPu8x','MnoWuwbzc3','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('0efbae62-6e9b-4c8d-961c-3b7b5130ddf8','KbGuU1Pmhj','EgUJCBhy5l','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('4f7f7c4b-7498-4beb-9f03-65d1d7b75bda','F9wtiRXkB7','kOegFhKFnw','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('8310abe0-ea3d-4b2f-856d-7b720c782e1d','2uTaAj0dWI','x9Wfdvlyqm','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);
INSERT INTO app.task VALUES ('1a169da6-4fd6-475c-851b-e7a5f1cc5b9f','ViqMZZtntZ','lNHobWvjKV','d3b5c915-9a78-4478-b487-fea43dfd74af',NULL,'2025-03-17 14:55:24',NULL);