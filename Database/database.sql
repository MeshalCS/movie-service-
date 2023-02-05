CREATE TABLE movies (
	id INTEGER PRIMARY KEY,
	Title TEXT NOT NULL,
	poster_path TEXT NOT NULL,
	description TEXT NOT NULL,
	genres TEXT NOT NULL
);
CREATE TABLE user_favorites (
	user_id INTEGER NOT NULL,
	movie_id INTEGER NOT NULL,
	PRIMARY KEY (user_id, movie_id),
	FOREIGN KEY (user_id) REFERENCES users (id),
	FOREIGN KEY (movie_id) REFERENCES movies (id)
);
