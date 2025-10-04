create table if not exists users (
    id SERIAL PRIMARY KEY,
    google_sub_id VARCHAR(50) UNIQUE,
    name varchar(255),
    given_name varchar(255),
    family_name varchar(255),
    picture varchar(255),
    email varchar(255),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);