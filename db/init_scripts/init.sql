CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS rankings (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    points_score INT NOT NULL,
    time_score INT NOT NULL
);

CREATE TABLE IF NOT EXISTS profiles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    picture VARCHAR(100) NOT NULL,
    description VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL
);

ALTER TABLE rankings
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE profiles
ADD CONSTRAINT fk_user_id
FOREIGN KEY (user_id) REFERENCES users(id);
