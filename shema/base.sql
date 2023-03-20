CREATE TABLE genre (
                       id SERIAL PRIMARY KEY,
                       title varchar(255)
);
CREATE TABLE author (
                        id SERIAL PRIMARY KEY,
                        Name varchar(255)
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

SELECT
    book.id,
    book.title,
    book.year,
    a.id,
    a.Name,
    g.id,
    g.title
FROM book
         INNER JOIN author a ON a.id = book.author_id
         INNER JOIN genre g ON g.id = book.genre_id;

SELECT
    book.id,
    book.title,
    book.year,
    g.id,
    g.title,
    a.id,
    a.Name
FROM book
         INNER JOIN genre g ON g.id = book.genre_id
         INNER JOIN author a on a.id = book.author_id
WHERE book.id = 1 LIMIT 1

SELECT id, name, created_at FROM "user"