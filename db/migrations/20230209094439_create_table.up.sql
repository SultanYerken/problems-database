CREATE TABLE IF NOT EXISTS problem 
(
    id SERIAL NOT NULL PRIMARY KEY,
    title VARCHAR(100) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    level VARCHAR (25) NOT NULL,
    topics TEXT NOT NULL,
    samples TEXT,
    created_at DATE NOT NULL,
    updated_at DATE 
);


CREATE TABLE IF NOT EXISTS topic 
(
    id SERIAL NOT NULL PRIMARY KEY,
    topic_name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS problem_topic 
(
    id SERIAL NOT NULL PRIMARY KEY,
    problem_id INT REFERENCES problem (id) ON DELETE CASCADE NOT NULL,
    topic_id INT REFERENCES topic (id) ON DELETE CASCADE NOT NULL
);


INSERT INTO topic (topic_name) VALUES ('Array');
INSERT INTO topic (topic_name) VALUES ('String');
INSERT INTO topic (topic_name) VALUES ('Hash Table');
INSERT INTO topic (topic_name) VALUES ('Dynamic Programming');
INSERT INTO topic (topic_name) VALUES ('Math');
INSERT INTO topic (topic_name) VALUES ('Sorting');
INSERT INTO topic (topic_name) VALUES ('Tree');
INSERT INTO topic (topic_name) VALUES ('Binary Search');
INSERT INTO topic (topic_name) VALUES ('Database');
INSERT INTO topic (topic_name) VALUES ('Other');
