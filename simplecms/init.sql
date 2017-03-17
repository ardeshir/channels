CREATE USER dbadmin WITH PASSWORD 'dbadmin';

CREATE DATABASE simplecms;

GRANT ALL PRIVILEGES ON DATABASE simplecms to dbadmin;

CREATE TABLE IF NOT EXISTS PAGES(
   id           SERIAL PRIMARY KEY,
   title        TEXT NOT NULL,
   content      TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS POSTS(
   id           SERIAL PRIMARY KEY,
   title        TEXT NOT NULL,
   content      TEXT NOT NULL,
   date_created DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS COMMENTS(
   id           SERIAL PRIMARY KEY,
   author       TEXT NOT NULL,
   comment      TEXT NOT NULL,
   date_created DATE NOT NULL,
   post_id      INT references POSTS(id)
);

