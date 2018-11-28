CREATE DATABASE world;
\connect world

CREATE TABLE countries(
    id          SERIAL          NOT NULL    PRIMARY KEY,
    name        VARCHAR(100)    NOT NULL,
    population  BIGINT
);

INSERT INTO countries (name, population) VALUES
    ('Russiano', 10000),
    ('Ukraine', 7000000000),
    ('Keks', 200),
    ('JoJoFANS', 1000000);
    
-- SELECT * FROM countries;