

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

CREATE SCHEMA library;

ALTER SCHEMA library OWNER TO dbuser;


CREATE TABLE library.book (
    id SERIAL PRIMARY KEY,
    writer varchar,
    title varchar,
    genre character varying,
    date date,
    isbn varchar,
    created_at timestamp without time zone,
    updated_at timestamp
);

ALTER TABLE library.book OWNER TO dbuser;

INSERT INTO library.book VALUES (1,'J.K. Rowling', 'Harry Potter and the Sorcerers Stone', 'Fantasy', '1997-06-26', '9780590353427', current_date, null);
INSERT INTO library.book VALUES (2,'Jeremy Clarkson', 'The World According to Clarkson', 'Non-Fiction', '2004-10-14', '9780316726055', current_date, null);