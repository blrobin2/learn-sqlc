CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    firstname VARCHAR NOT NULL,
    lastname VARCHAR NOT NULL
);

CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    user_id SERIAL NOT NULL REFERENCES users(id),
    task VARCHAR NOT NULL,
    done BOOLEAN NOT NULL
);
