CREATE TABLE IF NOT EXISTS movies (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	title VARCHAR NOT NULL,
	isbn VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS directors (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	first_name VARCHAR NOT NULL,
	last_name VARCHAR NOT NULL,
	movie_id INTEGER NOT NULL,
	FOREIGN KEY (movie_id) REFERENCES movies(id)
);

DELETE FROM movies;
DELETE FROM sqlite_sequence WHERE name = 'movies';

DELETE FROM directors;
DELETE FROM sqlite_sequence WHERE name = 'directors';

INSERT INTO movies (title, isbn) VALUES ('Movie One', '3829164');
INSERT INTO movies (title, isbn) VALUES ('Movie Two', '8372991');
INSERT INTO movies (title, isbn) VALUES ('Movie Three', '2818293');

INSERT INTO directors (first_name, last_name, movie_id) VALUES ('John', 'Smith', 1);
INSERT INTO directors (first_name, last_name, movie_id) VALUES ('Sarah', 'Geofferies', 2);
INSERT INTO directors (first_name, last_name, movie_id) VALUES ('Francis', 'Matthews', 3);
