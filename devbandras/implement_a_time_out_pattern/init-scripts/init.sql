-- 1. Felhasználó létrehozása (ha nem létezik)
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'app_user') THEN
        CREATE USER app_user WITH PASSWORD 'securepassword';
    END IF;
END $$;

-- 2. Séma létrehozása
CREATE SCHEMA IF NOT EXISTS app AUTHORIZATION app_user;

-- 3. Tábla létrehozása
CREATE TABLE IF NOT EXISTS app.book (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    published_year INT CHECK (published_year > 0),
    genre VARCHAR(100),
    price DECIMAL(10,2) NOT NULL
);

-- 4. Indexek létrehozása
CREATE INDEX IF NOT EXISTS idx_book_title ON app.book (title);
CREATE INDEX IF NOT EXISTS idx_book_author ON app.book (author);

-- 5. 30 rekord beszúrása
INSERT INTO app.book (title, author, published_year, genre, price) VALUES
    ('The Catcher in the Rye', 'J.D. Salinger', 1951, 'Fiction', 9.99),
    ('To Kill a Mockingbird', 'Harper Lee', 1960, 'Fiction', 12.50),
    ('1984', 'George Orwell', 1949, 'Dystopian', 14.99),
    ('Brave New World', 'Aldous Huxley', 1932, 'Science Fiction', 11.99),
    ('Fahrenheit 451', 'Ray Bradbury', 1953, 'Dystopian', 10.99),
    ('Moby-Dick', 'Herman Melville', 1851, 'Adventure', 13.50),
    ('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'Fiction', 8.99),
    ('War and Peace', 'Leo Tolstoy', 1869, 'Historical Fiction', 15.99),
    ('Pride and Prejudice', 'Jane Austen', 1813, 'Romance', 9.49),
    ('The Hobbit', 'J.R.R. Tolkien', 1937, 'Fantasy', 10.99),
    ('Crime and Punishment', 'Fyodor Dostoevsky', 1866, 'Psychological Fiction', 12.99),
    ('The Brothers Karamazov', 'Fyodor Dostoevsky', 1880, 'Philosophical', 14.50),
    ('Wuthering Heights', 'Emily Brontë', 1847, 'Gothic', 11.25),
    ('Dracula', 'Bram Stoker', 1897, 'Horror', 9.75),
    ('Frankenstein', 'Mary Shelley', 1818, 'Gothic', 10.50),
    ('Les Misérables', 'Victor Hugo', 1862, 'Historical Fiction', 13.99),
    ('The Count of Monte Cristo', 'Alexandre Dumas', 1844, 'Adventure', 14.99),
    ('One Hundred Years of Solitude', 'Gabriel García Márquez', 1967, 'Magical Realism', 12.50),
    ('The Picture of Dorian Gray', 'Oscar Wilde', 1890, 'Philosophical', 9.99),
    ('The Grapes of Wrath', 'John Steinbeck', 1939, 'Historical Fiction', 11.99),
    ('Catch-22', 'Joseph Heller', 1961, 'Satire', 10.99),
    ('Slaughterhouse-Five', 'Kurt Vonnegut', 1969, 'Science Fiction', 11.50),
    ('The Sun Also Rises', 'Ernest Hemingway', 1926, 'Fiction', 9.75),
    ('Lolita', 'Vladimir Nabokov', 1955, 'Psychological Fiction', 12.25),
    ('Beloved', 'Toni Morrison', 1987, 'Historical Fiction', 13.50),
    ('Gone with the Wind', 'Margaret Mitchell', 1936, 'Historical Fiction', 14.75),
    ('Dune', 'Frank Herbert', 1965, 'Science Fiction', 15.99),
    ('The Road', 'Cormac McCarthy', 2006, 'Post-Apocalyptic', 10.99),
    ('A Clockwork Orange', 'Anthony Burgess', 1962, 'Dystopian', 11.25),
    ('The Alchemist', 'Paulo Coelho', 1988, 'Fiction', 9.99);

-- 6. Jogosultságok beállítása az app_user számára
GRANT USAGE ON SCHEMA app TO app_user;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA app TO app_user;
GRANT ALL PRIVILEGES ON SEQUENCE app.book_id_seq TO app_user;