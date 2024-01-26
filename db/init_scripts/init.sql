CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS rankings (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    points_score INT NOT NULL
    time_score INT NOT NULL
);

CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    picture VARCHAR(100) NOT NULL,
    description VARCHAR(100) NOT NULL,
);

ALTER TABLE rankings
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE profile
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);
