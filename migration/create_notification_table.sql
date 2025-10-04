create table if not exists notifications (
 id serial primary key,
 type varchar(50),
 message varchar(255),
 owner_user_id int not null references users(id) on delete cascade,
 created_at TIMESTAMPTZ not null default (NOW() at TIME zone 'UTC')
);