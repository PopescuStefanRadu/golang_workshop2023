CREATE TABLE authors
(
    id   integer PRIMARY KEY,
    name text NOT NULL,
    bio  text
);

CREATE TABLE books
(
    id     integer PRIMARY KEY,
    author_id integer NOT NULL,
    title  text
);