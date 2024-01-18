CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS rankings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    score INT NOT NULL
);

ALTER TABLE rankings
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);
