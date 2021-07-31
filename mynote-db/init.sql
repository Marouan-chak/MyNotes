CREATE USER mynote;
CREATE DATABASE mynote;
GRANT ALL PRIVILEGES ON DATABASE mynote TO mynote;
CREATE TABLE notes (
	id serial PRIMARY KEY,
	title VARCHAR ( 50 )  NOT NULL,
	text VARCHAR ( 255 ) 
);