CREATE TABLE IF NOT EXISTS users (
    id INT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NULL
);

INSERT INTO users values (1, 'nolan', '2017-07-01T14:59:55');
INSERT INTO users values (2, 'test', '2017-07-01T15:59:55');
INSERT INTO users values (3, 'test', null);

CREATE TABLE IF NOT EXISTS test (
    id INT,
    num FLOAT DEFAULT NULL,
    b boolean DEFAULT NULL
);

INSERT INTO test values (1, 3.14, TRUE);
INSERT INTO test values (2, 4.014, FALSE);