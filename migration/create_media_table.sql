CREATE TABLE IF NOT EXISTS media (
	id SERIAL PRIMARY KEY,
	owner_user_id INT not null references "user"(id) on delete cascade,
	caption VARCHAR(255),
	file_storage_link VARCHAR(255),
	created_at TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')
);