

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(400) NOT NULL
);


CREATE UNIQUE INDEX idx_users_names ON users(name);

