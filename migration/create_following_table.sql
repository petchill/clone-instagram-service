create table if not exists following (
	id SERIAL primary key,
	user_id INT not null references "user"(id) on delete cascade,
	target_user_id INT not null references "user"(id) on delete cascade,
	created_at TIMESTAMPTZ not null default (NOW() at TIME zone 'UTC')
);