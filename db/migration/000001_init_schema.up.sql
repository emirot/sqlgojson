CREATE TABLE IF NOT EXISTS users (
    id INT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NULL
);

INSERT INTO users values (1, 'nolan', '2017-07-01T14:59:55.711Z');
INSERT INTO users values (2, 'test', '2017-07-01T15:59:55.711Z');
INSERT INTO users values (3, 'test', null);