create table if not exists following (
	id SERIAL primary key,
	user_id VARCHAR(50),
	target_user_id VARCHAR(50),
	created_at TIMESTAMPTZ not null default (NOW() at TIME zone 'UTC')
);