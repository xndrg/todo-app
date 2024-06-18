CREATE TABLE IF NOT EXISTS users (
	id serial NOT NULL UNIQUE,
	name varchar(255) NOT NULL,
	username varchar(255) NOT NULL UNIQUE,
	password_hash varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todo_lists (
	id serial NOT NULL UNIQUE,
	title varchar(255) NOT NULL,
	description varchar(255)
);

CREATE TABLE IF NOT EXISTS users_lists (
	id serial NOT NULL UNIQUE,
	user_id int REFERENCES users (id) ON DELETE CASCADE NOT NULL,
	list_id int REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE todo_items (
	id serial NOT NULL UNIQUE,
	title varchar(255) NOT NULL,
	description varchar(255),
	is_done boolean NOT NULL DEFAULT false
);

CREATE TABLE IF NOT EXISTS lists_items (
	id serial NOT NULL UNIQUE,
	item_id int REFERENCES todo_items (id) ON DELETE CASCADE NOT NULL,
	list_id int REFERENCES todo_lists (id) ON DELETE CASCADE NOT NULL
);

