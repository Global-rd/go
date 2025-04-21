SET timezone = 'Europe/Budapest';

CREATE USER HatakeKakashi WITH ENCRYPTED PASSWORD 'YvACldNZdAGtdUNwrz12uYihpixM2Wt2Ed0NRdtpC80kKXDnQjwQCvw5ZHVl1A1B';

GRANT ALL PRIVILEGES ON DATABASE books TO HatakeKakashi;


CREATE TABLE bookshelf (
    Id VARCHAR NOT NULL UNIQUE PRIMARY KEY,
    Title VARCHAR,
    Author VARCHAR,
    Published INT,
    Introduction VARCHAR,
    Price REAL,
    Stock INT
);


INSERT INTO bookshelf (Id, Title, Author, Published, Introduction, Price, Stock)
VALUES('TestID', 'TestTitle', 'TestAuthor', 2000, 'TestIntro', 123.45, 1);

