CREATE TABLE IF NOT EXISTS media (
	id SERIAL PRIMARY KEY,
	owner_user_id CHAR(20),
	caption CHAR(255),
	file_storage_link CHAR(255),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);