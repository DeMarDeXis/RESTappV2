CREATE TABLE users
(
    id serial not null unique,
    name VARCHAR(255) not null,
    username VARCHAR(255) not null unique,
    password_hash VARCHAR(255) not null
);

CREATE TABLE todo_lists
(
    id serial not null unique,
    title VARCHAR(255) not null,
    description VARCHAR(255)
);

CREATE TABLE users_lists
(
    id serial not null unique,
    user_id int REFERENCES users (id) on delete cascade not NULL,
    list_id int REFERENCES todo_lists (id) on delete cascade not NULL
);

CREATE TABLE todo_items
(
    id serial not null unique,
    title VARCHAR(255) not null,
    description VARCHAR(255),
    done boolean not null default false
);

CREATE TABLE lists_items
(
    id serial,
    item_id int REFERENCES todo_items (id) on delete cascade not NULL,
    list_id int REFERENCES todo_lists (id) on delete cascade not NULL
);