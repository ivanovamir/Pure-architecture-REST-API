CREATE TABLE genre (
    id SERIAL PRIMARY KEY,
    title varchar(255)
);
CREATE TABLE author (
    id SERIAL PRIMARY KEY,
    name varchar(255)
);
CREATE TABLE book (
    id SERIAL PRIMARY KEY,
    title varchar(255),
    year integer,
    genre_id integer references genre(id) on delete set null,
    author_id integer references author(id) on delete set null
);
CREATE TABLE user_book (
    user_id integer references "user"(id),
    book_id integer references book(id)
);
CREATE TABLE "user"(
    id SERIAL PRIMARY KEY,
    name varchar(255),
    created_at timestamptz
);
CREATE TABLE "token"(
    id SERIAL PRIMARY KEY,
    refresh_token varchar(32),
    user_id integer references "user"(id),
    created_at timestamptz
);