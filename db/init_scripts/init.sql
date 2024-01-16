CREATE TABLE IF NOT EXISTS player_ranking (
    player_id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    score INT NOT NULL
);
