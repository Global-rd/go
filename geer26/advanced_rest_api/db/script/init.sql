SET timezone = 'Europe/Budapest';

CREATE USER HatakeKakashi WITH ENCRYPTED PASSWORD 'YvACldNZdAGtdUNwrz12uYihpixM2Wt2Ed0NRdtpC80kKXDnQjwQCvw5ZHVl1A1B';

GRANT ALL PRIVILEGES ON DATABASE books TO HatakeKakashi;


CREATE TABLE bookshelf (
    Id VARCHAR NOT NULL UNIQUE PRIMARY KEY,
    Title VARCHAR NOT NULL,
    Author VARCHAR NOT NULL,
    Published INT NOT NULL,
    Introduction VARCHAR,
    Price REAL NOT NULL,
    Stock INT
);


INSERT INTO bookshelf
VALUES('TestID', 'TestTitle', 'TestAuthor', 2000, 'TestIntro', 123.45, 1);

