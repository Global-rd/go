-- noinspection SqlNoDataSourceInspectionForFile


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

CREATE SCHEMA movieapp;

ALTER SCHEMA movieapp OWNER TO movieadm;


CREATE TABLE movieapp.movies (
                                 id VARCHAR,
                                 title VARCHAR,
                                 release_date VARCHAR,
                                 imdb_id VARCHAR,
                                 director VARCHAR,
                                 writer VARCHAR,
                                 starts VARCHAR
);


ALTER TABLE movieapp.movies OWNER TO movieadm;

-- 20 Sample movie inserts for movieapp.movies table
INSERT INTO movieapp.movies VALUES ('1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n5o6p', 'The Shawshank Redemption', '1994-09-23', 'tt0111161', 'Frank Darabont', 'Stephen King, Frank Darabont', 'Tim Robbins');
INSERT INTO movieapp.movies VALUES ('2b3c4d5e-6f7g-8h9i-0j1k-2l3m4n5o6p7q', 'The Godfather', '1972-03-24', 'tt0068646', 'Francis Ford Coppola', 'Mario Puzo, Francis Ford Coppola', 'Marlon Brando');
INSERT INTO movieapp.movies VALUES ('3c4d5e6f-7g8h-9i0j-1k2l-3m4n5o6p7q8r', 'The Dark Knight', '2008-07-18', 'tt0468569', 'Christopher Nolan', 'Jonathan Nolan, Christopher Nolan', 'Christian Bale');
INSERT INTO movieapp.movies VALUES ('4d5e6f7g-8h9i-0j1k-2l3m-4n5o6p7q8r9s', 'Pulp Fiction', '1994-10-14', 'tt0110912', 'Quentin Tarantino', 'Quentin Tarantino', 'John Travolta');
INSERT INTO movieapp.movies VALUES ('5e6f7g8h-9i0j-1k2l-3m4n-5o6p7q8r9s0t', 'Inception', '2010-07-16', 'tt1375666', 'Christopher Nolan', 'Christopher Nolan', 'Leonardo DiCaprio');
INSERT INTO movieapp.movies VALUES ('6f7g8h9i-0j1k-2l3m-4n5o-6p7q8r9s0t1u', 'The Matrix', '1999-03-31', 'tt0133093', 'Lana Wachowski, Lilly Wachowski', 'Lana Wachowski, Lilly Wachowski', 'Keanu Reeves');
INSERT INTO movieapp.movies VALUES ('7g8h9i0j-1k2l-3m4n-5o6p-7q8r9s0t1u2v', 'Goodfellas', '1990-09-19', 'tt0099685', 'Martin Scorsese', 'Nicholas Pileggi, Martin Scorsese', 'Ray Liotta');
INSERT INTO movieapp.movies VALUES ('8h9i0j1k-2l3m-4n5o-6p7q-8r9s0t1u2v3w', 'Fight Club', '1999-10-15', 'tt0137523', 'David Fincher', 'Chuck Palahniuk, Jim Uhls', 'Brad Pitt');
INSERT INTO movieapp.movies VALUES ('9i0j1k2l-3m4n-5o6p-7q8r-9s0t1u2v3w4x', 'The Lord of the Rings: The Fellowship of the Ring', '2001-12-19', 'tt0120737', 'Peter Jackson', 'J.R.R. Tolkien, Fran Walsh, Philippa Boyens', 'Elijah Wood');
INSERT INTO movieapp.movies VALUES ('0j1k2l3m-4n5o-6p7q-8r9s-0t1u2v3w4x5y', 'Forrest Gump', '1994-07-06', 'tt0109830', 'Robert Zemeckis', 'Winston Groom, Eric Roth', 'Tom Hanks');
INSERT INTO movieapp.movies VALUES ('1k2l3m4n-5o6p-7q8r-9s0t-1u2v3w4x5y6z', 'The Silence of the Lambs', '1991-02-14', 'tt0102926', 'Jonathan Demme', 'Thomas Harris, Ted Tally', 'Jodie Foster');
INSERT INTO movieapp.movies VALUES ('2l3m4n5o-6p7q-8r9s-0t1u-2v3w4x5y6z7a', 'Interstellar', '2014-11-07', 'tt0816692', 'Christopher Nolan', 'Jonathan Nolan, Christopher Nolan', 'Matthew McConaughey');
INSERT INTO movieapp.movies VALUES ('3m4n5o6p-7q8r-9s0t-1u2v-3w4x5y6z7a8b', 'The Departed', '2006-10-06', 'tt0407887', 'Martin Scorsese', 'William Monahan, Alan Mak, Felix Chong', 'Leonardo DiCaprio');
INSERT INTO movieapp.movies VALUES ('4n5o6p7q-8r9s-0t1u-2v3w-4x5y6z7a8b9c', 'Gladiator', '2000-05-05', 'tt0172495', 'Ridley Scott', 'David Franzoni, John Logan, William Nicholson', 'Russell Crowe');
INSERT INTO movieapp.movies VALUES ('5o6p7q8r-9s0t-1u2v-3w4x-5y6z7a8b9c0d', 'The Green Mile', '1999-12-10', 'tt0120689', 'Frank Darabont', 'Stephen King, Frank Darabont', 'Tom Hanks');
INSERT INTO movieapp.movies VALUES ('6p7q8r9s-0t1u-2v3w-4x5y-6z7a8b9c0d1e', 'Saving Private Ryan', '1998-07-24', 'tt0120815', 'Steven Spielberg', 'Robert Rodat', 'Tom Hanks');
INSERT INTO movieapp.movies VALUES ('7q8r9s0t-1u2v-3w4x-5y6z-7a8b9c0d1e2f', 'Schindler''s List', '1993-12-15', 'tt0108052', 'Steven Spielberg', 'Thomas Keneally, Steven Zaillian', 'Liam Neeson');
INSERT INTO movieapp.movies VALUES ('8r9s0t1u-2v3w-4x5y-6z7a-8b9c0d1e2f3g', 'Parasite', '2019-10-11', 'tt6751668', 'Bong Joon Ho', 'Bong Joon Ho, Jin Won Han', 'Song Kang-ho');
INSERT INTO movieapp.movies VALUES ('9s0t1u2v-3w4x-5y6z-7a8b-9c0d1e2f3g4h', 'Whiplash', '2014-10-10', 'tt2582802', 'Damien Chazelle', 'Damien Chazelle', 'Miles Teller');
INSERT INTO movieapp.movies VALUES ('0t1u2v3w-4x5y-6z7a-8b9c-0d1e2f3g4h5i', 'The Prestige', '2006-10-20', 'tt0482571', 'Christopher Nolan', 'Jonathan Nolan, Christopher Nolan', 'Hugh Jackman');