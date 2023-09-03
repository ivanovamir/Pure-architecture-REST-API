BEGIN;
CREATE TABLE IF NOT EXISTS "user" (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL UNIQUE,
    password_hash BYTEA
);

CREATE UNIQUE INDEX idx_user_phone ON "user"(phone);
CREATE UNIQUE INDEX idx_user_email ON "user"(email);

CREATE TABLE IF NOT EXISTS author (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT NOT NULL,
    is_alive BOOL,
    date_of_birth DATE,
    date_of_death DATE CHECK ( date_of_death <= date(now()) )
);

CREATE TABLE IF NOT EXISTS book (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    number_of_pages INT,
    author_id INT REFERENCES author(id) ON DELETE SET NULL,
    publication_date DATE CHECK ( publication_date <= date(now()) )
);

-- Вставка записи №1: Агата Кристи
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Агата', 'Кристи', 'Мэри', false, TO_DATE('15/09/1890', 'DD/MM/YYYY'), TO_DATE('12/01/1976', 'DD/MM/YYYY'));

-- Вставка записи №2: Уильям Шекспир
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Уильям', 'Шекспир', '', false, TO_DATE('26/04/1564', 'DD/MM/YYYY'), TO_DATE('23/04/1616', 'DD/MM/YYYY'));

-- Вставка записи №3: Джейн Остин
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Джейн', 'Остин', '', false, TO_DATE('16/12/1775', 'DD/MM/YYYY'), TO_DATE('18/07/1817', 'DD/MM/YYYY'));

-- Вставка записи №4: Фёдор Достоевский
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Фёдор', 'Достоевский', 'Михайлович', false, TO_DATE('11/11/1821', 'DD/MM/YYYY'), TO_DATE('09/02/1881', 'DD/MM/YYYY'));

-- Вставка записи №5: Джордж Оруэлл
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Джордж', 'Оруэлл', '', false, TO_DATE('25/06/1903', 'DD/MM/YYYY'), TO_DATE('21/01/1950', 'DD/MM/YYYY'));

-- Вставка записи №6: Лев Толстой
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Лев', 'Толстой', 'Николаевич', false, TO_DATE('09/09/1828', 'DD/MM/YYYY'), TO_DATE('20/11/1910', 'DD/MM/YYYY'));

-- Вставка записи №7: Дж. К. Роулинг
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Джоан', 'Роулинг', 'Кэтлин', true, TO_DATE('31/07/1965', 'DD/MM/YYYY'), null);

-- Вставка записи №8: Джордж Р. Р. Мартин
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Джордж', 'Мартин', 'Рэймонд Ричард', true, TO_DATE('20/09/1948', 'DD/MM/YYYY'), null);

-- Вставка записи №9: Эрнест Хемингуэй
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Эрнест', 'Хемингуэй', 'Миллер', false, TO_DATE('21/07/1899', 'DD/MM/YYYY'), TO_DATE('02/07/1961', 'DD/MM/YYYY'));

-- Вставка записи №10: Джон Стейнбек
INSERT INTO author (name, surname, patronymic, is_alive, date_of_birth, date_of_death)
VALUES ('Джон', 'Стейнбек', 'Эрнест', false, TO_DATE('27/02/1902', 'DD/MM/YYYY'), TO_DATE('20/12/1968', 'DD/MM/YYYY'));

COMMIT;